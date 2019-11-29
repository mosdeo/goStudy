package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("into ServeHTTP(w http.ResponseWriter, r *http.Request)")
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9080", mux)
}
