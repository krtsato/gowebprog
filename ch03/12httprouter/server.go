// リスト3.12

package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter" // go get github.com/julienschmidt/httprouter
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()

	// /hello/hoge とすると, 完全一致せずとも /hello にルーティングする
	// /hoge とすると, 完全一致しないと正しくルーティングされない
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
