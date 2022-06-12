package middleware

import (
	"net/http"
	"prival-api/pkg/token"

	"github.com/gin-gonic/gin"
)

var (
	authHeader = "x-auth"
	authCtxKey = "auth"
)

func Auth(token token.Token) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		base64Str := ctx.GetHeader(authHeader)
		if len(base64Str) < 1 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "authenticate headers are missing",
			})
			return
		}

		payload, err := token.Verify(base64Str)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Set(authCtxKey, payload.UserID)
		ctx.Next()
	}
}

func GetAuthKey() string {
	return authCtxKey
}
