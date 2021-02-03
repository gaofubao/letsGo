package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"io"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	// 创建一个tracer，但是返回一个no-op tracer
	//tracer := opentracing.GlobalTracer()

	// 初始化一个jaeger tracer
	tracer, closer := Init("hello-world")
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
	// 定义该函数自己的span
	//span := rootSpan.Tracer().StartSpan(
	//	"formatString",				// 操作名称
	//	opentracing.ChildOf(rootSpan.Context()),	// 建立因果关系
	//	)
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(log.String("event", "string-format"), log.String("value", helloStr))
	return helloStr
}

func printHello(ctx context.Context, helloStr string) {
	//span := rootSpan.Tracer().StartSpan(
	//	"printHello",
	//	opentracing.ChildOf(rootSpan.Context()),
	//	)
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	println(helloStr)
	span.LogKV("event", "println")
}