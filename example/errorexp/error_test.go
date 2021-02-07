package errorexp

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func TestError(t *testing.T)  {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
}

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func TestEcho(t *testing.T)  {
	for _, req := range []string{"", "hello"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}

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

// 怎样判断一个错误值具体代表的是哪一类错误？

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func printError(i int, err error)  {
	if err == nil {
		fmt.Println("nil error")
		return
	}
	err = underlyingError(err)
	switch err {
	case os.ErrClosed:
		fmt.Printf("error(closed)[%d]: %s\n", i, err)
	case os.ErrInvalid:
		fmt.Printf("error(invalid)[%d]: %s\n", i, err)
	case os.ErrPermission:
		fmt.Printf("error(permission)[%d]: %s\n", i, err)
	}
}


