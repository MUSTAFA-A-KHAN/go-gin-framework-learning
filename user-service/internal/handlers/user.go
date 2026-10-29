package handlers

import (
	"net/http"
	"strconv"

	"github.com/MUSTAFA-A-KHAN/go-gin-framework-learning/user-service/internal/models"
	"github.com/gin-gonic/gin"
)

var U = []models.User{{
	Name:      "Mustafa",
	Age:       450,
	Address:   "Pune",
	IsPremium: true,
}, {
	Name:      "Arthur Miller",
	Age:       26,
	Address:   "NY",
	IsPremium: false,
},
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PONG",
	})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": U,
	})
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"message": U[id],
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	U[id].Name = "NewName"
	c.JSON(200, map[string]string{"Updated": U[id].Name})
}
