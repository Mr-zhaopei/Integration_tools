package logger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//定义日志级别结构体信息

//自己定义一个日志文件库
//基于uint16构建一个日志级别的类型
type LogLevel uint16

const (
	//日志级别定义
	UNKNOWN LogLevel = iota
	DEBUG
	// TRACE
	INFO
	WARNGIN
	ERROR
	FATAL
)

//parseLogLevel日志级别解析
func parseLogLevel(s string) (LogLevel, error) {
	//将接收到的字符串转换为大写
	s = strings.ToLower(s)

	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNGIN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		// err := fmt.Errorf()返回一个格式化输出错误
		err := errors.New("无效的日志Level")
		return UNKNOWN, err
	}

}

//getInfo(runtime.Caller)可以获取到调用时的代码文件路径、行数等信息，在打印日志时常常使用
func getInfo(skip int) (funcName, fileName string, line int) {
	pc, filepath, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed error\n")
	}

	//函数名称获取
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	//打印日志文件另获取
	fileName = path.Base(filepath)
	return
}
