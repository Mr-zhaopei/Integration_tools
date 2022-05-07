package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//向文件中写日志相关代码
type Filelogger struct {
	Level       LogLevel //日志等级
	filePath1   string   //文件路径
	fileName    string   //日志文件名
	fileObj     *os.File //日志文件指针
	errFileObj  *os.File //错误日志文件指针
	maxFileSize int64    //日志文件大小，超过该大小之后日志文件需要进行切割
}

//定义一个创建日志文件的函数
//levelStr日志等级，fp日志文件路径，fn，文件名称，max maxSize日志文件大小
func NewFilelogger(levelStr, fp, fn string, maxSize int64) *Filelogger {
	//将传递进来的日志等级进行转换为字符串大写
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &Filelogger{
		Level:       logLevel,
		filePath1:   fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}

	err = f1.initFile() //按照文件路径和文件名称打开
	if err != nil {
		panic(err)
	}
	return f1
}

//打开文件
func (f *Filelogger) initFile() error {

	//打开普通文件
	fullFileName := path.Join(f.filePath1, f.fileName) //路径拼接
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}

	//构建一个错误日志文件
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}

	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil

}

//日志文件关闭
func (f *Filelogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *Filelogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

//定义个判断日志文件大小的方法，返回一个bool值
func (f *Filelogger) checksize(file *os.File) bool {
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info faild err:%v\n", err)
		return false
	}
	//若果日志文件大于最大文件大小，则返回一个true进行日志文件切割
	return fileinfo.Size() > f.maxFileSize
}

func (f *Filelogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割
	nowstr := time.Now().Format("20060102150405000")
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath1, fileinfo.Name()) //拿到当前的日志文件完整路径
	newLogname := fmt.Sprintf("%s.bak%s", logName, nowstr)
	//1.关闭当前的日志文件
	file.Close()
	//2.备份以下日志文件 rename xxx.log xxx.log.bak201908091709
	os.Rename(logName, newLogname) //文件重命名
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	//4.将打开的新的日志文件对象赋值给f.fileObj
	return fileObj, nil

}

func (f *Filelogger) log(lv LogLevel, format string, a ...interface{}) { //该处借鉴fmt.Println()函数进行输出格式
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) //格式化字符串信息
		//获取当前时间
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3) //skip根据调用getInfo函数来进行判断执行文件的行号
		if f.checksize(f.fileObj) {
			newfile, err := f.splitFile(f.fileObj) //日志文件
			if err != nil {
				return
			}
			f.fileObj = newfile
		}
		fmt.Fprintf(f.fileObj, "[%s][%s] [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checksize(f.errFileObj) {
				newfile, err := f.splitFile(f.errFileObj) //日志文件
				if err != nil {
					return
				}
				f.errFileObj = newfile
			}
			//如果要记录的日志大于等于ERROR级别，我还需要在err日志文件中记录一遍
			fmt.Fprintf(f.errFileObj, "[%s][%s] [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

//日志返回的方法
func (f *Filelogger) Debug(format string, a ...interface{}) {
	if f.enable(DEBUG) {
		f.log(DEBUG, format, a...)
	}
}
func (f *Filelogger) Warning(format string, a ...interface{}) {
	f.log(WARNGIN, format, a...)
}

func (f *Filelogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *Filelogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *Filelogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
