package middlewares

import (
	"gin-user-management/lib"
	"gin-user-management/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	service services.AuthService
	logger  lib.Logger
}

func NewJWTMiddleware(
	service services.AuthService,
	logger lib.Logger) JWTMiddleware {
	return JWTMiddleware{
		service: service,
		logger:  logger,
	}
}

func (jwtMiddleware JWTMiddleware) Setup() {}

func (jwtMiddleware JWTMiddleware) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")

		splitHeader := strings.Split(authHeader, " ")

		if len(splitHeader) > 1 {
			token := splitHeader[1]

			auth, err, id := jwtMiddleware.service.Authorize(token)
			if auth {
				context.Set("userId", id)
				context.Next()
				return
			}
			//todo AbortWithStatusJSON
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			jwtMiddleware.logger.Zap.Error(err)
			context.Abort()
			return
		}
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		context.Abort()
	}
}
