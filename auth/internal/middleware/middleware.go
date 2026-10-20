package middleware

import (
	"fmt"
	"go/auth-service/internal/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentification() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		clientToken := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.UserType)

		c.Next()
	}
}

type TokenRequest struct {
	Token string `json:"token"`
}

func AdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenReq TokenRequest
		if err := c.ShouldBindJSON(&tokenReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		token := tokenReq.Token
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			return
		}

		claims, err := helpers.ValidateToken(token)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_type": claims.UserType})
	}
}

func TakeIdFromToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenReq TokenRequest
		if err := c.ShouldBindJSON(&tokenReq); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response on user-service"})
			return
		}
		token := tokenReq.Token
		claims, err := helpers.ValidateToken(token)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		fmt.Println("Taken id:", claims.Uid)
		c.JSON(http.StatusOK, gin.H{"id": claims.Uid, "email": claims.Email})
	}
}
