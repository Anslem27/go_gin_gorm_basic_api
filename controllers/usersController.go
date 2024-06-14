package controllers

import (
	"go_crud/initializers"
	"go_crud/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-crypt/x/bcrypt"
	"github.com/golang-jwt/jwt"
)

func SignUpUser(c *gin.Context) {
	// get user password - emil off req body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if (c.Bind(&body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Request body parameters are missing or invalid",
		})
		return
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to hash password",
		})
		return
	}
	user := models.User{Email: body.Email, Password: string(hash)}

	// create a new user

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   result.Error.Error(),
		})
		return
	}

	// return a new user

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User created successfully",
	})
}

func LogIn(c *gin.Context) {
	// get email and password off req body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if (c.Bind(&body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Request body parameters are missing or invalid",
		})
		return
	}
	// look up user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "User not found",
		})
		return

	}
	// compare password and hashed
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid password",
		})
		return
	}

	//generate jwd token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                                    // subject
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // expiration time
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Failed to generate log in token",
		})
		return
	}
	// `c.SetSameSite(http.SameSiteLaxMode)` is setting the SameSite attribute of the cookie to Lax mode
	// for the current request context in the Gin framework. This attribute is used for preventing certain
	// types of cross-site request forgery (CSRF) attacks by restricting when cookies are sent in a
	// cross-origin request. In Lax mode, cookies are sent with top-level navigations and when submitting
	// forms using methods other than GET.
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30 /* month */, "", "", false, true)
	// send it back
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		// "token":   tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	// user(models.User).Email

	// if exists {
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
	// }
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to get users",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"users":   users,
	})
}
