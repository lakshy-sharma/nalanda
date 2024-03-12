package controllers

import (
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/models"
	"nalanda_backend/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	appConfig *initializers.AppConfig
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
	DB        *gorm.DB
}

func NewAuthController(applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) AuthController {
	return AuthController{
		appConfig: applicationConfig,
		InfoLog:   infoLogger,
		ErrorLog:  errorLogger,
		DB:        dbConn,
	}
}

func (ac *AuthController) UserSignUp(ctx *gin.Context) {
	// Parse and bind the data to a model.
	var payload *models.UserSignUpInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": ""})
		return
	}

	newUser := models.User{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PasswordHash: hashedPassword,
		Role:         "user",
		Verified:     true,
	}

	result := ac.DB.Create(&newUser)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that mail already exists"})
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": "Something went wrong"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "User Created"})
}

func (ac *AuthController) UserSignIn(ctx *gin.Context) {
	// Parse the input.
	var payload models.UserLoginInput
	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Perform validation.
	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Email or Password"})
		return
	}
	if err := utils.VerifyPassword(user.PasswordHash, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Email or Password"})
		return
	}

	// Generate tokens for the user.
	accessToken, err := utils.CreateToken(ac.appConfig.Backend.AccessTokenAge, user.ID, ac.appConfig.Backend.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	refresh_token, err := utils.CreateToken(ac.appConfig.Backend.RefreshTokenAge, user.ID, ac.appConfig.Backend.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Set the tokens in the cookie.
	ctx.SetCookie("access_token", accessToken, int(ac.appConfig.Backend.AccessTokenAge*time.Duration(60*time.Second)), "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, int(ac.appConfig.Backend.RefreshTokenAge*time.Duration(60*time.Second)), "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", int(ac.appConfig.Backend.AccessTokenAge*time.Duration(60*time.Second)), "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
