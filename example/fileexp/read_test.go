package fileexp

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// ioutil.ReadFile()
func Read1() string {
	f, err := ioutil.ReadFile("/home/aishu/access.log")
	if err != nil {
		fmt.Println("read file failed", err)
		return ""
	}
	return string(f)
}

func TestRead1(t *testing.T) {
	fmt.Println(Read1())
}

func Read2() string {
	f, err := os.Open("/home/aishu/access.log")
	if err != nil {
		fmt.Println("read file failed", err)
		return ""
	}
	defer f.Close()

	var chunk []byte
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf failed", err)
			return ""
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}
	return string(chunk)
}

func TestRead2(t *testing.T) {
	fmt.Println(Read2())
}

func Read3() string {
	f, err := os.Open("/home/aishu/access.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	var chunks []byte
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf...)
	}
	return string(chunks)
}

func TestRead3(t *testing.T)  {
	fmt.Println(Read3())
}

func Read4() string {
	f, err := os.Open("/home/aishu/access.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(r)
}

func TestRead4(t *testing.T) {
	fmt.Println(Read4())
}

