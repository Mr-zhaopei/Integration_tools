package logger

import (
	"fmt"
	"time"
)

//该文件功能实现的是一个向终端进行输出的日志简单库

//consuleLogger为调用的时候显示的日志级别
type consulLogger struct {
	Level LogLevel
}

func NewconsulLogger(levelStr string) *consulLogger {
	//将用户端传入的levelStr的字符串进行级别解析
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &consulLogger{
		Level: level,
	}
}

//判断传入的日志等级是否存在
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case WARNGIN:
		return "WARNGIN"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "错误日志等级不存在"
	}

}

//日志级别的判断：主要是为了打印日志界别范围的一个判断
func (c consulLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c consulLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...) //格式化字符串信息
		//获取当前时间
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3) //skip根据调用getInfo函数来进行判断执行文件的行号
		fmt.Printf("[%s][%s] [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c consulLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c consulLogger) Warning(format string, a ...interface{}) {
	c.log(WARNGIN, format, a...)
}

func (c consulLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c consulLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c consulLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
