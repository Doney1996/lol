package common

import (
	"github.com/gin-gonic/gin"
)

var gameType = map[string]struct{}{
	"class": struct{}{},
	"clone": struct{}{},
	"chaos": struct{}{},
}

// Biz 存放游戏的类型 没有则是class
func Biz() gin.HandlerFunc {
	return func(c *gin.Context) {
		typeName := c.Request.Header.Get("type")
		_, isPresent := gameType[typeName]
		if isPresent {
			c.Set("gameType", typeName)
		} else {
			c.Set("gameType", "class")
		}
	}
}
