package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"lol/controller"
	"net/http"
	"runtime/debug"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(Recover)
	// 告诉gin框架去哪里找模板文件
	r.Static("/static", "./static")
	v1Group := r.Group("/v1")
	{
		v1Group.POST("/addRecord", controller.AddRecord)
		v1Group.POST("/disable", controller.DisableHero)
		v1Group.POST("/enable", controller.EnableAllHero)
		v1Group.POST("/addHero", controller.AddHero)
		v1Group.GET("/getAllHero", controller.GetAllHero)
		v1Group.POST("/jiesuan", controller.JieSuan)
	}
	return r
}

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
				"code": "1",
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
