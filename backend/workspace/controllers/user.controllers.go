package controllers

import (
	"fmt"
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/models"
	"nalanda_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	appConfig   *initializers.AppConfig
	DB          *gorm.DB
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewUserController(applicationConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) UserController {
	return UserController{
		appConfig:   applicationConfig,
		DB:          dbConn,
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}
}

func (uc *UserController) GetUserData(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	userResponse := &models.UserDataResponse{
		ID:        currentUser.ID,
		Email:     currentUser.Email,
		Role:      currentUser.Role,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var payload *models.UserUpdateInput
	var err error
	currentUser := ctx.MustGet("currentUser").(models.User)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	currentUser.Email = payload.Email
	currentUser.FirstName = payload.FirstName
	currentUser.LastName = payload.LastName
	currentUser.PasswordHash, err = utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSon
		fmt.Println("")
	}
	uc.DB.Save(&currentUser)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "action": "Update", "message": "Updated User Email"})
}

func (uc *UserController) ResetUserPassword(ctx *gin.Context) {
	var payload *models.UserPasswordResetInput
	currentUser := ctx.MustGet("currentUser").(models.User)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Check if old and new passwords are the same.
	if payload.OldPassword == payload.NewPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Previous and Current Passwords Cannot be Same!"})
		return
	}
	// Validate the old password.
	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.PasswordHash), []byte(payload.OldPassword)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid Old Password. Validation Failed."})
		return
	}

	// Hash the new password and save it inside the database.
	hashedPassword, err := utils.HashPassword(payload.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	currentUser.PasswordHash = hashedPassword
	uc.DB.Save(&currentUser)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "action": "Update", "message": "Updated User Password"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	// Perform a soft delete of the user.
	uc.DB.Delete(&currentUser)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"id": currentUser.ID, "action": "DELETE"}})
}
