package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserLogin - ..
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/logins", func(c *gin.Context) {
		var u UserLogin
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": u,
		})
	})
	r.Run(":8080")
}
