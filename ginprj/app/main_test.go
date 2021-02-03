package main

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T)  {
	header := []byte("{\"Accept\":[\"*/*\"],\"Content-Length\":[\"0\"],\"User-Agent\":[\"curl/7.29.0\"],\"X-B3-Sampled\":[\"1\"],\"X-B3-Spanid\":[\"965ea2eed5255c6c\"],\"X-B3-Traceid\":[\"f843a66b719db2a6965ea2eed5255c6c\"],\"X-Forwarded-Proto\":[\"http\"],\"X-Request-Id\":[\"23610f2d-d405-99d5-a636-6c54b7b7bcdd\"]}")
	for k, v := range header {
		fmt.Println(k, v)
	}
}
