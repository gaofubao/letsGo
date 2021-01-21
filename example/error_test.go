package example

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

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


