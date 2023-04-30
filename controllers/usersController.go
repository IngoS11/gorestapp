package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IngoS11/gorestapp/initializers"
	"github.com/IngoS11/gorestapp/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup godoc
//
//	@Summary		add a user
//	@Description	add a user with password to the system
//	@Tags			users
//	@Accept			json
//	@Produc			json
//	@Param			user	body		controllers.User	true	"Add model"
//	@Success		200		{object}	controllers.User
//	@Router			/users [post]
func Signup(c *gin.Context) {
	// Get the email/password of the request body
	var body User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to ready body",
		})

		return
	}

	// Hash the password before storing it into the db
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	fmt.Println(hash)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// create the user in the database
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add User to the database",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Login godoc
//
//	@Summary		login user
//	@Description	login user with email and password
//	@Tags			users
//	@Accept			json
//	@Produc			json
//	@Param			user	body		controllers.User	true	"Add model"
//	@Success		200		{object}	controllers.User
//	@Router			/users/login [post]
func Login(c *gin.Context) {
	// get the email and pass off request body
	var body User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password not specified in body",
		})

		return
	}

	// Lookup user in database
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid email or password",
		})

		return

	}

	// Generate a JWT Token and return it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 240).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// send back token as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*240, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

// Validate godoc
//
//	@Summary		validate user
//	@Description	validate user via his jwt token in cookie
//	@Tags			users
//	@Accept			json
//	@Produc			json
//	@Success		200	{object}	controllers.User
//	@Router			/users/validate [get]
func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User is valid",
	})
}
