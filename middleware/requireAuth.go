package middleware

import (
	"fmt"
	"go_crud/initializers"
	"go_crud/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Validating user checks.")
	// get the token from cookie from the header
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(401, gin.H{"success": false, "error": "Unauthorized"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	// validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// our secret key
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// find user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach to request
		c.Set("user", user)

		// proceed
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
