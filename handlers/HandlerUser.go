package handlers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "helix/dgsi/api/models"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

// get all users
func (handler UserHandler) Index(c *gin.Context) {
	users := []m.User{}	
	handler.db.Table("tbl_users").Order("id desc").Find(&users)
	c.JSON(http.StatusOK, &users)
}


