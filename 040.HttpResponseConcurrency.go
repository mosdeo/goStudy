package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func main() {
	var accmulateRequest = 0
	var PrintNumGoroutine = func(w http.ResponseWriter, r *http.Request) {
		go func() {

			accmulateRequest++
			msg := fmt.Sprintf("ID:%v, 累積Request:%d, Num of Goroutine:%d", GetGID(), accmulateRequest, runtime.NumGoroutine())
			<-time.After(time.Second)
			fmt.Fprintf(w, msg)
			fmt.Println(msg)
		}()
	}

	http.HandleFunc("/concurrencyExam", PrintNumGoroutine)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
