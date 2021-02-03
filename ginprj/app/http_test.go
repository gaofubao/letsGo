package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

var dep = os.Getenv("DEP")
var port = os.Getenv("SE_PORT")
var headers = []string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",
	"x-ot-span-context",
	"x-ot-span-context",
	"x-datadog-trace-id",
	"x-datadog-parent-id",
	"x-datadog-sampled",
}

func foo(w http.ResponseWriter, r *http.Request) {
	head := r.Header
	io.WriteString(w, "foo header:\n")

	res, err := http.NewRequest("GET", "http://"+dep+"/api/sheet/bar", nil)
	fmt.Println("access: " + "http://" + dep + "/api/sheet/bar")
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("content-type", "application/json")

	client := &http.Client{}

	for k, i := range head {
		io.WriteString(w, k+": "+strings.Join(i, ", ")+"\n")
		// for _, v := range i {
		//  res.Header.Add(k, v)
		// }
	}
	fmt.Fprintln(w)

	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("content-type", "application/json")

	for _, k := range headers {
		v := head.Get(k)
		if v != "" {
			res.Header.Set(k, head.Get(k))
			// fmt.Println(k, v)
		}

	}
	resp, _ := client.Do(res)
	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)

		return
	}

	_, err = w.Write(result)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, "\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	head := r.Header
	io.WriteString(w, "bar header:\n")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("content-type", "application/json")
	for k, i := range head {
		io.WriteString(w, k+": "+strings.Join(i, ", ")+"\n")
	}

}

func TestNet(t *testing.T) {
	fmt.Println("start .....")
	for i, k := range headers {
		head := strings.Title(k)
		headers[i] = head
	}
	//fmt.Println(headers)
	http.HandleFunc("/api/sheet/foo", foo)
	http.HandleFunc("/api/sheet/bar", bar)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
