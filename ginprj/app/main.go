package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/app1", HandleApp1)
	r.GET("/app2", HandleApp2)
	r.GET("/header", HandlePing)

	r.Run()
}

func HandlePing(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HandleApp1(c *gin.Context) {
	reqHeader := c.Request.Header

	headerJson, _ := json.Marshal(reqHeader)
	WriteLog(string(headerJson))

	//go ConnApp("135.31.227.92", 5)
	ConnApp("135.31.227.92", reqHeader,5)
	c.Writer.WriteString(string(headerJson))
}

func HandleApp2(c *gin.Context) {
	reqHeader := c.Request.Header

	headerJson, _ := json.Marshal(reqHeader)
	WriteLog(string(headerJson))

	go ConnApp("127.0.0.1", reqHeader,2)
	c.Writer.WriteString(string(headerJson))
}

func WriteLog(s string) {
	f, err := os.OpenFile("/root/access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, fmt.Sprintf(s))
	w.Flush()
}

func ConnApp(ip string, header map[string][]string, t int) {
	time.Sleep(time.Duration(t)*time.Second)

	//resp, err := http.Get("http://"+ip+":8080/header")
	req, _ := http.NewRequest("GET", "http://"+ip+":8080/header", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-Id", header["X-Request-Id"][0])
	req.Header.Set("X-B3-Traceid", header["X-B3-Traceid"][0])
	req.Header.Set("X-B3-Spanid", header["X-B3-Spanid"][0])
	req.Header.Set("X-Forwarded-Proto", header["X-Forwarded-Proto"][0])
	req.Header.Set("X-B3-Sampled", header["X-B3-Sampled"][0])

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	headerJson, _ := json.Marshal(resp.Header)
	WriteLog(string(headerJson))
}


/*
   "X-B3-Sampled": [
       "1"
   ],
   "X-B3-Spanid": [
       "f8ee3f8a2998cdc5"
   ],
   "X-B3-Traceid": [
       "7a54a3f140847716f8ee3f8a2998cdc5"
   ],
   "X-Forwarded-Proto": [
       "http"
   ],
   "X-Request-Id": [
       "7ef5873a-21ce-97f3-9c7d-9661c6b020c8"
   ]
 */
