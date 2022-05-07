package logger

//定义logger接口
type Logger interface {
	Debug(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}
