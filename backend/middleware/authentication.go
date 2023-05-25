package middleware

import (
	"log"
	"net/http"
	"pawAPIbackend/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
*
Logica de validação fica noiddleware
*/
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "failed to process request - token not found",
			})
			return
		}
		token, err := service.ValidateToken(authHeader)
		if !token.Valid {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token is not valid",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		//deixa que o a rota fique protegida e passe para o controlador
		//book.POST("/", middleware.Authorized(), controller.InsertBook)
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
