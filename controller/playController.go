package controller

import (
	"lol/common"
	"lol/entity"
	"lol/repository/play"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string
	Password string
}

type LoginRes struct {
	Token  string        `json:"token"`
	Player entity.Player `json:"player"`
}

// RegisterUser  注册用户
func RegisterUser(c *gin.Context) {
	var player entity.Player
	err := c.BindJSON(&player)
	common.DealErr(err)
	play.AddPlayer(&player)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "注册成功！",
	})
}

// Login 登录
func Login(c *gin.Context) {
	var loginReq LoginReq
	err := c.BindJSON(&loginReq)
	common.DealErr(err)

	isPass, player := play.LoginCheck(loginReq.Username, loginReq.Password)
	if isPass {
		generateToken(c, player)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "验证失败," + err.Error(),
		})
	}

}

// 生成令牌
func generateToken(c *gin.Context, player *entity.Player) {
	j := &common.JWT{
		SigningKey: []byte("ruozheyouxi"),
	}
	claims := common.CustomClaims{
		ID:       player.Id,
		Name:     player.Name,
		GameName: player.GameName,
		UserName: player.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "ruozheyouxi",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	data := LoginRes{
		Player: *player,
		Token:  token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*common.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}
