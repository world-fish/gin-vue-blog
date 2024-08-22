package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitCore()
	global.DB = core.InitGorm()
	fmt.Println(global.DB)

}
