package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TempUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := CustomClaims{
			ID:             1,
			Name:           "欧东林",
			GameName:       "碧血葬香魂",
			UserName:       "oudonglin",
			StandardClaims: jwt.StandardClaims{},
		}
		c.Set("claims", claims)
	}
}
