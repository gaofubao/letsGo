package httpexp

import (
	"net/http"
	"testing"
)

/*
命名类型
- ConnState
- Dir
- HandlerFunc
- Header
- SameSite
结构体
- Client
- Cookie
- ProtocolError
- PushOptions
- Request
- Response
- ServeMux
- Server
- Transport
接口
- CloseNotifier
- CookieJar
- File
- FileSystem
- Flusher
- Handler
- Hijacker
- Pusher
- ResponseWriter
- RoundTripper
 */

func SayHello(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Hello"))
}

func TestServer(t *testing.T)  {
	http.HandleFunc("/hello", SayHello)

	http.ListenAndServe(":8081", nil)
}
