package fileexp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func Write1() {
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("file open error", err)
		return
	}

	//fmt.Println(ioutil.ReadAll(f))

	n, err := io.WriteString(f, "test")
	if err != nil {
		fmt.Println("write error", err)
		return
	}
	fmt.Println("bytes", n)
}

func TestWrite1(t *testing.T)  {
	Write1()
}

func Write2() {
	f, err := os.OpenFile("access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, fmt.Sprintf("test1"))
	w.Flush()
}

func TestWrite2(t *testing.T)  {
	Write2()
}
