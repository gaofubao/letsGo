package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"letsGo/ginprj/tutorial"
	"log"
	"net/http"
)

func main() {
	tracer, closer := tutorial.InitTracer("publisher")
	defer closer.Close()

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("publish", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		helloStr := r.FormValue("helloStr")
		println(helloStr)
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
