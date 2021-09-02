package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"lol/common"
	"lol/controller"
	"lol/entity"
	"lol/expection"
	"net/http"
	"runtime/debug"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(Recover)
	r.Use(Cors())
	//r.Use(common.JWTAuth())
	r.Use(common.TempUser())
	r.Use(common.Biz())
	PathRouter(r)

	return r
}

// PathRouter 添加路由的路径
func PathRouter(r *gin.Engine) {

	r.POST("/login", controller.Login)
	r.POST("/register", controller.RegisterUser)

	v1Group := r.Group("/v1")
	{
		// 对局
		v1Group.POST("/addRecord", controller.AddRecord)

		//英雄
		v1Group.GET("/getAllHero", controller.GetAllHero)
		v1Group.GET("/getHeroBySeason", controller.GetHeroBySeason)

		v1Group.POST("/disable", controller.DisableHero)
		v1Group.POST("/enable", controller.EnableAllHero)
		v1Group.POST("/enableById", controller.EnableAllHeroById)
		v1Group.POST("/addHero", controller.AddHero)
		v1Group.GET("/getAllHeroList", controller.GetAllHeroInfoList)
		v1Group.GET("/recent", controller.GetRecentResult)

		//赛季
		v1Group.POST("/openNewSeason", controller.OpenNewSeason)
		v1Group.POST("/closeSeason", controller.CloseSeason)

		//对局
		v1Group.POST("/openNewMatch", controller.OpenNewMatch)
		v1Group.POST("/getLastMatch", controller.GetLastMatch)
		v1Group.POST("/closeNewMatch", controller.CloseNewMatch)

	}
}

// Cors 解决跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// Recover 全局异常处理
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替

			c.JSON(http.StatusOK, entity.Result{
				Code:    errorToBizErr(r).Code,
				Message: errorToBizErr(r).Msg,
				Data:    nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToBizErr(r interface{}) expection.BizErr {
	switch v := r.(type) {
	case expection.BizErr:
		return v
	case error:
		return expection.BizErr{
			Code: 500,
			Msg:  v.Error(),
		}
	default:
		return expection.BizErr{
			Code: 501,
			Msg:  "系统繁忙",
		}
	}
}
