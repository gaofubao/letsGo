package goroutineexp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

// Keep yourself busy or do the work yourself
func TestGoroutine1(t *testing.T)  {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Gorotine")
	}
	http.HandleFunc("/", helloHandler)
	// 无法获得goroutine的运行状态
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	// 空的select将永远阻塞
	select {}
}

// Never start a goroutine without knowning when it will stop
func TestGoroutine2(t *testing.T)  {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Goroutine")
	}
	http.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestGoroutine3(t *testing.T)  {
	hellohandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Goroutine")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hellohandler)
	// 无法感知该goroutine什么时候停止
	go http.ListenAndServe(":8081", http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func serverApp4() {
	hellohandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Goroutine")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hellohandler)

	http.ListenAndServe(":8080", mux)
}

func ServerDebug4()  {
	http.ListenAndServe(":8081", http.DefaultServeMux)
}

func TestGoroutine4(t *testing.T) {
	go ServerDebug4()
	serverApp4()
}

func serverApp5() {
	hellohandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Goroutine")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hellohandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func ServerDebug5()  {
	// log.fatal()调用了os.Exit()，一般只用在main()或init()中
	log.Fatal(http.ListenAndServe(":8081", http.DefaultServeMux))
}

func TestGoroutine5(t *testing.T) {
	go ServerDebug5()
	go serverApp5()
	select {}
}

// 1.goroutine由调用者来决定是否在后台执行
// 2.可以感知groutine什么时候停止
// 3.可以控制grutine的退出
//func ServerApp6(addr string, handler http.Handler, stop <-chan struct{}) error {
//	s := http.Server{
//		Addr: addr,
//		Handler: handler,
//	}
//	go func() {
//		<-stop
//		s.Shutdown(context.Background())
//	}()
//	return s.ListenAndServe()
//}

//func ServerDebug6(addr string, handler http.Handler, stop <-chan struct{}) error {
//	s := http.Server{
//		Addr: addr,
//		Handler: handler,
//	}
//	go func() {
//		<-stop
//		s.Shutdown(context.Background())
//	}()
//	return s.ListenAndServe()
//}
//
//func TestServer6(t *testing.T) {
//	hellohandler := func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "Hello Goroutine")
//	}
//
//	done := make(chan error, 2)
//	stop := make(chan struct{})
//
//	go func() {
//		done <- ServerDebug6(":8081", hellohandler, stop)
//	}()
//
//	go func() {
//		done <- ServerApp6(":8080", hellohandler, stop)
//	}()
//
//	var stopped bool
//	for i := 0; i < cap(done); i++ {
//		if err := <-done; err != nil {
//			fmt.Println("error: %v", err)
//		}
//		if !stopped {
//			stopped = true
//			close(stop)
//		}
//	}
//}

