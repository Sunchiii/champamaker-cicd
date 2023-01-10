package middleware

import (
	"net/http"
	"strings"

	"github.com/Sunchiii/champamker-service/helpers"
	"github.com/Sunchiii/champamker-service/responses"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if err := helpers.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, responses.Status{
				Status:  http.StatusUnauthorized,
				Massage: "UNAUTHORIZRED",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}
	}
}
