package middleware

import (
	"fmt"
	"log"
	"nalanda_backend/initializers"
	"nalanda_backend/models"
	"nalanda_backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthenticateUser(appConfig *initializers.AppConfig, dbConn *gorm.DB, infoLogger *log.Logger, errorLogger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Fetch the access token from the cookies andthe auth header (Bearer field).
		var accessToken string
		cookie, err := ctx.Cookie("access_token")
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}
		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		// Validate the token provided by the user.
		userID, err := utils.ValidateToken(accessToken, appConfig.Backend.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		// Get the user whose token we have validated.
		// Note: Soft deleted records are ignored automatically.
		var user models.User
		result := dbConn.First(&user, "id = ?", fmt.Sprint(userID))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token doesnt exists"})
			return
		}

		// Set the user in the context to be used by our handlers.
		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
