package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ContentTypeIsJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there is a content type, for now just take JSON
		if c.ContentType() != "" && c.ContentType() != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
			c.Abort()
			return
		}

		// If passed check, continue to next handler
		c.Next()
	}
}

func AuthVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}

		// Parse: takes key func to verify (plus any additional code), return back public key
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// TODO: can check signing type with the key type used
			return PublicKey, nil
		})
		if token == nil || !token.Valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Get claims: get username and check expiration time
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			expTime := claims["exp"].(float64)
			if expTime < float64(time.Now().Unix()) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": "expired token"})
				c.Abort()
				return
			}

			if !claims.VerifyIssuer("astroauth-server", true) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": "invalid issuer"})
				c.Abort()
				return
			}

			c.Set("username", claims["sub"])
		}

		// Token validated, can continue onto next protected routes
		c.Next()
	}
}
