package main

import (
	"fmt"
	"lol/routers"
)

func main() {
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
