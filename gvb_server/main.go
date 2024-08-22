package main

import (
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitCore()

	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("嘻嘻嘻")
	global.Log.Errorln("嘻嘻嘻")
	global.Log.Infoln("嘻嘻嘻")

	logrus.Warnln("嘻嘻嘻")
	logrus.Errorln("嘻嘻嘻")
	logrus.Infoln("嘻嘻嘻")

	// 连接数据库
	global.DB = core.InitGorm()

}
