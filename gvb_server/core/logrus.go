package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"os"
	"path"
)

// ANSI 颜色代码
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// LogFormatter 是一个自定义的 Logrus 日志格式化器
type LogFormatter struct{}

// Format 方法定义了日志的输出格式
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志等级 设定对应的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 设置缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 从全局配置中获取日志前缀
	logPrefix := global.Config.Logger.Prefix
	// 格式化时间戳
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	//判断日志条目包是否含调用者信息
	if entry.HasCaller() {
		// 自定义文件路径
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line) //文件名
		funcVal := entry.Caller.Function                                                 // 函数名

		// 自定义输出格式
		// \x1b[%dm[%s]\x1b[0m  这是 ANSI 转义序列, 用于设置终端的文本颜色, %d是颜色代码
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s \n", logPrefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		// 如果日志条目不包含调用者信息
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s \n", logPrefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// InitLogger 初始化一个 Logrus 日志记录器
func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                // 实例化
	mLog.SetOutput(os.Stdout)                           // 设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                  // 设置自定义的Formatter 自动调用其Format方法

	// 从全局配置中解析日志等级
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 设置最低日志级别

	// 对全局logrus进行同样配置
	InitDefaultLogger()

	// todo: 设置输出到按日期命名的文件
	//logFilePath := getLogFilePath()
	//file, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	//mLog.Out = file

	return mLog
}

// InitDefaultLogger 全局默认的logrus配置
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低日志级别
}

//func getLogFilePath() string {
//	// 指定日志文件存放目录
//	logDir := "logs"
//
//	// 创建或确认目录存在
//	if err := os.MkdirAll(logDir, 0o755); err != nil {
//		panic(fmt.Sprintf("Failed to create log directory: %v", err))
//	}
//
//	// 生成以当前日期命名的日志文件名
//	t := time.Now()
//	fileName := fmt.Sprintf("%s_%s.log", t.Format("2006-01-02"), "app")
//
//	// 返回完整文件路径
//	return filepath.Join(logDir, fileName)
//}
