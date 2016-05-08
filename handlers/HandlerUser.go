package handlers

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"io"
	"errors"
	"strconv"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	jwt_lib "github.com/dgrijalva/jwt-go"
	m "helix/dgsi/api/models"
	"helix/dgsi/api/config"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

//get all users
func (handler UserHandler) Index(c *gin.Context) {
	users := []m.User{}	
	handler.db.Table("tbl_user").Order("id desc").Find(&users)
	c.JSON(http.StatusOK, &users)
}

//create new user
func (handler UserHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	username := c.PostForm("username")
	password := c.PostForm("password")
	companyid := c.PostForm("company_id")

	if (username == "") {
		respond(http.StatusBadRequest,"Please supply the user's username",c,true)
	} else if (password == "") {
		respond(http.StatusBadRequest,"Please supply the user's password",c,true)
	} else if (companyid == "") {
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
			    //DECRYPTION
			    // result, err := decrypt(key, ciphertext)
			    // if err != nil {
			    //     log.Fatal(err)
			    // }
			    // fmt.Printf("%s\n", result)
			} else {
				year,_ := strconv.Atoi(fmt.Sprintf("%s",user.Clientid))
				clientid = strconv.Itoa(year + 1)
			}

		    toEncrypt := []byte(password)
		    ciphertext,_ := encrypt([]byte(config.GetString("CRYPT_KEY")), toEncrypt)

		    encryptedPassword := fmt.Sprintf("%0x", ciphertext)
			result := handler.db.Exec("INSERT INTO tbl_user VALUES(null,?,?,?,?,?,?,?)",clientid,username,encryptedPassword,companyid,now,now,"active")

			if (result.RowsAffected == 1) {
				c.JSON(http.StatusCreated, generateJWT(clientid))
			} else {
				respond(http.StatusBadRequest,"Unable to create new user, Please try again",c,true)
			}
		}
	}
}

func generateJWT(clientid string) m.AuthenticatedUser {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims["ID"] = "clientid"
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString([]byte(config.GetString("TOKEN_KEY")))
	user := m.AuthenticatedUser{}
	user.Token = tokenString
    return user
}

func encrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    b := base64.StdEncoding.EncodeToString(text)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
    return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    if len(text) < aes.BlockSize {
        return nil, errors.New("ciphertext too short")
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return nil, err
    }
    return data, nil
}

