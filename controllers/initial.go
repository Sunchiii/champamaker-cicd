package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initial() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"data": "has running!",
		})
	}
}
