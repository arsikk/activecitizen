package middleware

import (
	"ActiveCitizenRESTAPI/helper"
	"github.com/gin-gonic/gin"
)

func RoleCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := helper.GetTokenFromRequest(context)

	}
}
