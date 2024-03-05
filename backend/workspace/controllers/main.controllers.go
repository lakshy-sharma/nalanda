package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(ctx *gin.Context) {
	message := "Welcome to Nalanda: An Open Society for Academics"
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func Index(ctx *gin.Context) {
	message := "Welcome to Nalanda: An Open Society for Academics"
	ctx.JSON(http.StatusOK, gin.H{"message": message, "status": "success"})
}
