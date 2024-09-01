package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"time"
)

func InitGorm() *gorm.DB {
	// 检查全局配置中的数据库主机是否为空
	if global.Config.Mysql.Host == "" {
		global.Log.Printf("未配置数据库信息")
		return nil
	}

	// 生成数据库的 DSN (Data Source Name) 字符串，用于连接数据库
	dsn := global.Config.Mysql.Dsn()
	global.Log.Printf("查看数据库连接地址：%s", dsn)

	// 定义一个变量 mysqlLogger，用于存储 Gorm 的日志接口
	var mysqlLogger logger.Interface
	// 根据系统环境配置日志级别
	if global.Config.System.Env == "debug" {
		// 如果系统环境是开发模式 ("debug")，则设置为显示所有 SQL 语句的详细信息
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 否则，在生产环境 ("release") 中，只记录错误信息
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	// 使用 Gorm 的 Open 方法连接数据库，传入 DSN 和配置参数
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 配置日志记录器
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatal(fmt.Sprintf("数据库连接失败: %s", err))
	}

	// 获取通用数据库对象 sql.DB，以便进行更细粒度的控制
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)               // 设置连接池中的最大闲置连接数
	sqlDB.SetMaxOpenConns(100)              // 设置连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置连接池最大生存时间，不能超过mysql的wait_timeout

	// 返回配置好的 Gorm 数据库实例
	return db
}