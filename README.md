# Integration_tools 该工具是一个go的大集合工具

## 【1】Logger：用来实现一个日志格式化终端输出以及格式化文件输出的作用
### 该日志文件的简单使用如下：
```go
package main

import (
	"day03/mylogger"
)
//测试我们自己写的日志库
//如果想要所有的目录下全部都是用该log，可以在全局进行声明
var log mylogger.Logger
func main() {
	//log = mylogger.NewFilelogger("info", "./", "service.log", 10*1024*1024) //终端日志实例
	log = mylogger.NewConsolelog("info")                                    //文件日志实例
	for {
		id, name := 100, "test"
		log.Debug("这是一条Debug的日志id=%d,name=%s,", id, name)
		log.Info("这是一条info日志id=%d,name=%s,", id, name)
		log.Warning("这是一跳Warning的日志id=%d,name=%s,", id, name)
		log.Error("这是一跳Warning的日志id=%d,name=%s,", id, name)
		log.Fatal("这是一跳Fatal的日志id=%d,name=%s,", id, name)
		// time.Sleep(time.Second)
	}

}
```
