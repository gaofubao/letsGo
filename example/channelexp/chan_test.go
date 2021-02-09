package channelexp

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T)  {
	// channel 类型
	// 有缓存channel 无缓存channel
	ch1 := make(chan string, 10)	// 双向channel
	fmt.Println(ch1)

	ch2 := make(chan<- float64)	// 只可接收数据
	fmt.Println(ch2)

	ch3 := make(<-chan int)		// 只可发送数据
	fmt.Println(ch3)

	// 检查channel是否关闭
	ch1<- "abc"			// 发送数据
	v, ok := <-ch1		// 接收数据
	fmt.Println(v, ok)

	close(ch1)
	v1, ok1 := <-ch1
	fmt.Println(v1, ok1)

	c := make(chan int)
	defer close(c)
	go func() {
		c<- 3 + 4
	}()
	i := <- c
	fmt.Println(i)
}

func worker(done chan bool)  {
	time.Sleep(time.Second)
	done<- true
}

func TestSync(t *testing.T)  {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}


