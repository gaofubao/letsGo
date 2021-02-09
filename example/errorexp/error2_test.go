package errorexp

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// 错误
// 在Go语言中，错误是可以预期的，并且不是非常严重，不会影响程序的运行。对于这类问题，可以用返回错误给调用者的方法，让调用者自己决定如何处理。
// error接口
// 一般，error接口用于当方法或函数执行遇到错误时进行返回，而且是最后一个返回值
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("math: square root of negative number")
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return f, nil
}

func TestSqrt(t *testing.T)  {
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)
}

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

type temp interface {
	Temp() bool
}

// Opaque errors
func TestOpaque(t *testing.T)  {
	_, err := os.Open("access.txt")
	_, ok := err.(temp)
	fmt.Println(ok)
}

// error断言

// error 嵌套
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open file failed")
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read file failed")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	//return config, errors.Wrap(err, "could not read config")
	return config, errors.WithMessage(err, "could not read config")
}

func TestReadfile(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		//fmt.Println(err)
		fmt.Printf("original error: %T %v\n", errors.WithStack(err), errors.WithStack(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		//os.Exit(1)
	}
}

type MyError struct {
	when time.Time
	what string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.when, e.what)
}

func oops() error {
	return MyError{
		time.Date(2021, 2, 8, 13, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func TestOops(t *testing.T)  {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}

// As
func TestAs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError

		if errors.As(err, &pathError) {
			fmt.Println("Failed at path: ", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

// Is
func TestIs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}

// New
func TestNew(t *testing.T)  {
	//err := errors.New("emit macho dwarf: elf header corrupted")

	const name, id = "tom", 20
	err := fmt.Errorf("user %q (id %d) not found", name, id)

	if err != nil {
		fmt.Print(err)
	}
}

// Unwrap
// func Unwrap(err error) error

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

