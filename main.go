package main

import (
	"github.com/riba2534/OnlineJudge/conf"
	"github.com/riba2534/OnlineJudge/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
