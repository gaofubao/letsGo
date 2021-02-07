package errorexp

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

// 错误
// 在Go语言中，错误是可以预期的，并且不是非常严重，不会影响程序的运行。对于这类问题，可以用返回错误给调用者的方法，让调用者自己决定如何处理。
// error接口
// 一般，error接口用于当方法或函数执行遇到错误时进行返回，而且是最后一个返回值

// Sentinel Error
func TestSentinel(t *testing.T)  {
	r := bytes.NewReader([]byte("abc"))
	_, err := r.Read(make([]byte, 10))
	if err == io.EOF {
		log.Fatal("read failed: ", err)
	}
}

// Error type
func TestErrortype(t *testing.T)  {
	//os.PathError{}
}

// Opaque errors
func TestOpaque(t *testing.T)  {

}

// error断言

// error 嵌套

// panic
// Go语言是一门静态的强类型语言，很多问题都尽可能地在编译时捕获，但是有一些只能在运行时检查，比如数组越界访问、不相同的类型强制转换等，这类运行时的问题会引起panic异常
// 此外，我们自己也可以抛出panic异常
// panic异常是一种非常严重的情况，会让程序中断运行，是程序崩溃，所以如果是不影响程序运行的错误，不要使用panic，使用普通错误error即可。

// recover
// 通常情况下，我们不对panic异常做任何处理，因为既然它影响程序运行的异常，就让它直接崩溃即可。
// 但是也有一些特例，比如程序崩溃前做一些资源释放的处理，这时就需要从paic异常中恢复，才能完成处理。
func TestRecover(t *testing.T)  {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("", "root", "123456")
}

func connectMySQL(ip, username, password string) {
	if ip == "" {
		panic("ip不能为空")
	}
}

