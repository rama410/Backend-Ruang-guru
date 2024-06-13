package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		tokenString, err := ctx.Cookie("session_token")
		header := ctx.GetHeader("Content-Type")
		if header == "application/json" {
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "missing token"})
				return
			}
		}else if tokenString == ""{
			ctx.JSON(http.StatusSeeOther, model.ErrorResponse{Error: "cookie empty"})
			return
		}

		
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
			return
		}
		
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized"})
			return
		}
		
		
		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
