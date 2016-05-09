package handlers

import (
	"net/http"
	"fmt"
	"time"
	"io"
	"strconv"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	jwt_lib "github.com/dgrijalva/jwt-go"
	m "helix/dgsi/api/models"
	"helix/dgsi/api/config"
	"gopkg.in/redis.v3"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

//get all users
func (handler UserHandler) Index(c *gin.Context) {
	if IsTokenValid(c) {
		users := []m.User{}	
		handler.db.Table("tbl_user").Order("id desc").Find(&users)
		fmt.Println("HEADER ---> " + c.Request.Header.Get("Authorization"))
		c.JSON(http.StatusOK, &users)
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)	
	}
}

//create new user
func (handler UserHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	username := c.PostForm("username")
	password := c.PostForm("password")
	companyid := c.PostForm("company_id")

	if (strings.TrimSpace(username) == "") {
		respond(http.StatusBadRequest,"Please supply the user's username",c,true)
	} else if (strings.TrimSpace(password) == "") {
		respond(http.StatusBadRequest,"Please supply the user's password",c,true)
	} else if (strings.TrimSpace(companyid) == "") {
		respond(http.StatusBadRequest,"Please supply the user's company id",c,true)
	} else {
		//check if username already existing
		user := m.User{}	
		handler.db.Table("tbl_user").Where("username = ?",username).Find(&user)

		if (user.Clientid != "") {
			respond(http.StatusBadRequest,"Username already taken",c,true)
		} else {
			//check count of users
			user := m.User{}	
			handler.db.Table("tbl_user").Last(&user)
			var clientid = ""

			//check if there are no users yet
			if user.Clientid == "" {
				year := strconv.Itoa(now.Year())
				clientid = year + "0000001"
			    //TODO DECRYPTION
			    // result, err := decrypt(key, ciphertext)
			} else {
				year,_ := strconv.Atoi(fmt.Sprintf("%s",user.Clientid))
				clientid = strconv.Itoa(year + 1)
			}

		    encryptedPassword := encrypt([]byte(config.GetString("CRYPT_KEY")), password)

			result := handler.db.Exec("INSERT INTO tbl_user VALUES(null,?,?,?,?,?,?,?)",clientid,username,encryptedPassword,companyid,now,now,"active")

			if (result.RowsAffected == 1) {
				c.JSON(http.StatusCreated, generateJWT(clientid))
			} else {
				respond(http.StatusBadRequest,"Unable to create new user, Please try again",c,true)
			}
		}
	}
}

//user authentication
func (handler UserHandler) Auth(c *gin.Context) {
	if IsTokenValid(c) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if (strings.TrimSpace(username) == "") {
			respond(http.StatusBadRequest,"Please supply the user's username",c,true)
		} else if (strings.TrimSpace(password) == "") {
			respond(http.StatusBadRequest,"Please supply the user's password",c,true)
		} else {
			//check if username already existing
			user := m.User{}	
			handler.db.Table("tbl_user").Where("username = ?",username).Find(&user)

			if user.Clientid == "" {
				respond(http.StatusBadRequest,"Account not found!",c,true)
			} else {
				decryptedPassword := decrypt([]byte(config.GetString("CRYPT_KEY")), user.Password)
				//invalid password
				if decryptedPassword != password {
					respond(http.StatusBadRequest,"Account not found!",c,true)
				} else {
					//authentication successful
					authenticatedUser := m.AuthenticatedUser{}
					authenticatedUser.Id = user.Id
					authenticatedUser.Clientid = user.Clientid
					authenticatedUser.Username = user.Username
					authenticatedUser.Companyid = user.Companyid
					authenticatedUser.Token = generateJWT(user.Clientid).Token
					c.JSON(http.StatusOK, authenticatedUser)
				}					
			}
		}
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)	
	}
}

//update user
func (handler UserHandler) Update(c *gin.Context) {
	if IsTokenValid(c) {
		client_id := c.Param("client_id")
		username := c.PostForm("username")
		companyid := c.PostForm("company_id")

		if (strings.TrimSpace(username) == "") {
			respond(http.StatusBadRequest,"Please supply the user's username",c,true)
		} else if (strings.TrimSpace(companyid) == "") {
			respond(http.StatusBadRequest,"Please supply the user's company id",c,true)
		} else {
			//check if user is existing based on the passed client id
			currentUser := m.User{}	
			handler.db.Table("tbl_user").Where("clientid = ?",client_id).Find(&currentUser)

			if (currentUser.Clientid == "") {
				respond(http.StatusBadRequest,"User record not found",c,true)
			} else {
				//check if username already existing
				otherUser := m.User{}	
				handler.db.Table("tbl_user").Where("clientid != ? AND username = ?",client_id, username).Find(&otherUser)

				if (otherUser.Clientid != "") {
					respond(http.StatusBadRequest,"Username already taken",c,true)
				} else {
					if currentUser.Username == username && currentUser.Companyid == companyid {
						respond(http.StatusBadRequest,"No changes detected",c,true)
					} else {
						now := time.Now().UTC()
						result := handler.db.Exec("UPDATE tbl_user SET username = ?, companyid = ?, date_updated = ? WHERE clientid = ?",username,companyid,now,client_id)
						if (result.RowsAffected == 1) {
							updatedUser := m.User{}
							handler.db.Table("tbl_user").Where("clientid = ?",client_id).Find(&updatedUser)
							c.JSON(http.StatusOK, updatedUser)
						} else {
							respond(http.StatusBadRequest,"Failed to update user record",c,true)
						}
					}
				}
			}
		}
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)
	}
}

//logout
func (userHandler UserHandler) Logout(c *gin.Context) {
	if IsTokenValid(c) {
		username := c.PostForm("username")
		if (strings.TrimSpace(username) == "") {
			respond(http.StatusBadRequest,"Please supply the user's username",c,true)
		} else {
			//add token to blacklist
			AddTokenToRedis(c)
			respond(http.StatusOK,"Successfully logged out from the system",c,false)
		}		
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)
	}
}

func generateJWT(clientid string) m.JWT {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims["ID"] = "clientid"
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString([]byte(config.GetString("TOKEN_KEY")))
	user := m.JWT{}
	user.Token = tokenString
    return user
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func AddTokenToRedis(c *gin.Context) {
    client := redis.NewClient(&redis.Options{
        Addr:     ":6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    token := c.Request.Header.Get("Authorization")
    err := client.Set(token, token, time.Duration(86400)*time.Second).Err()
    if err != nil {
        panic(err)
    } else {
    	fmt.Println("Successfully written in redis")
    	result, err := client.Get(token).Result()
    	if (err == nil) {
    		fmt.Println("RESULT ---> " + result)
    	}
    }
    defer client.Close()
}

func IsTokenValid(c *gin.Context) bool {
    client := redis.NewClient(&redis.Options{
        Addr:     ":6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    defer client.Close()
    token := c.Request.Header.Get("Authorization")
    result, _ := client.Get(token).Result()
	if (result != "") {
		return false
	}
	return true
}


