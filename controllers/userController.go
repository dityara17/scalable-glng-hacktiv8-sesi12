package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-jwt-glng-aditya/database"
	"go-jwt-glng-aditya/helpers"
	"go-jwt-glng-aditya/models"
	"log"
	"net/http"
)

var(
	appJSON = "application/json"
)

func UserRegister(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_,_ = db,contentType
	User := models.User{}

	if contentType == appJSON{
		err := c.ShouldBindJSON(&User)
		if err != nil {
			log.Fatalln(err)
		}
	}else{
		fmt.Println("masuk sini")
		err := c.ShouldBind(&User)
		if err != nil {
			log.Fatalln(err)
		}
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": 			User.ID,
		"email": 		User.Email,
		"full_name": 	User.FullName,
	})

}

func UserLogin(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_,_ = db,contentType
	User := models.User{}
	password := ""

	if contentType == appJSON{
		err := c.ShouldBindJSON(&User)
		if err != nil {
			log.Fatalln(err)
		}
	}else{
		fmt.Println("masuk sini")
		err := c.ShouldBind(&User)
		if err != nil {
			log.Fatalln(err)
		}
	}

	password = User.Password



	err := db.Debug().Where("email=?",User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized",
			"message" : "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password),[]byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized,gin.H{
			"error" : "Unauthorized",
			"message" : "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID,User.Email)

	c.JSON(http.StatusOK,gin.H{
		"token" : token,
	})

}
