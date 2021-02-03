package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"io/ioutil"
	"letsGo/ginprj/tutorial"
	"net/http"
	"net/url"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	// 初始化一个jaeger tracer
	tracer, closer := tutorial.InitTracer("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	helloTo := os.Args[1]

	// 定义一个span，并添加span标签
	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)
	span.Finish()

	// context.Context可以用来传播上下文
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// 记录Logs，有两种方式span.LogFields()、span.LogKV()
	helloStr := formatString(ctx, helloTo)
	printHello(ctx, helloStr)
}


func formatString(ctx context.Context, helloTo string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloTo", helloTo)
	url := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := Do(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err.Error())
	}

	helloStr := string(resp)
	span.LogFields(log.String("event", "string-format"), log.String("value", helloStr))
	return helloStr
}

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloStr", helloStr)
	url := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	// 注入tracer
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	if _, err := Do(req); err != nil {
		ext.LogError(span, err)
		panic(err.Error())
	}
}

func Do(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	return body, nil
}