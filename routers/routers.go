package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"lol/common"
	"lol/controller"
	"net/http"
	"runtime/debug"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.Use(Cors())

	r.POST("/login", controller.Login)
	r.POST("/register", controller.RegisterUser)

	//r.Use(common.JWTAuth())
	r.Use(common.TempUser())
	PathRouter(r)

	r.Use(Recover)
	return r
}

// PathRouter 添加路由的路径
func PathRouter(r *gin.Engine) {
	v1Group := r.Group("/v1")
	{
		v1Group.POST("/addRecord", controller.AddRecord)
		v1Group.POST("/disable", controller.DisableHero)
		v1Group.POST("/enable", controller.EnableAllHero)
		v1Group.POST("/enableById", controller.EnableAllHeroById)
		v1Group.POST("/addHero", controller.AddHero)
		v1Group.GET("/getAllHero", controller.GetAllHeroTop)
		v1Group.GET("/getAllHeroList", controller.GetAllHeroInfoList)
		v1Group.POST("/jiesuan", controller.JieSuan)
		v1Group.GET("/recent", controller.GetRecentResult)

		//赛季
		v1Group.POST("/openNewSeason", controller.OpenNewSeason)
		v1Group.POST("/closeSeason", controller.CloseSeason)

		//hero
		v1Group.GET("/disableHeroIds", controller.GetHeroBySeason)

		//对局
		v1Group.GET("/openNewMatch", controller.OpenNewMatch)

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
			c.JSON(http.StatusOK, gin.H{
				"code": "500",
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
