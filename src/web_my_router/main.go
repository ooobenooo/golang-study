package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

/**
实现了Handler接口，自定义路由
*/
func (mux *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloRouter(w, r)
		return
	}

	http.NotFound(w, r)
}

func sayHelloRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello my router.")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
