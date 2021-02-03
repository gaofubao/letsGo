package logexp

import (
	"io"
	"log"
	"os"
	"testing"
)

var (
	Info 	*log.Logger
	Warning *log.Logger
	Error 	*log.Logger
)

func init() {
	errFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	// 定义日志输出形式、日志前缀、日志属性
	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[Warning] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func TestLog(t *testing.T)  {
	Info.Println("Tom accessed web")
	Warning.Println("The traffic is too large")
	Error.Println("The instance was down")
}

// 在标准库中什么时候使用函数，什么时候使用方法？
