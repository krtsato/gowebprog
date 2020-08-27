package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handlerは指定しない -> DefaultServeMuxをハンドラとして利用
	}

	// 関数をハンドラに変換して、DefaultServeMux に登録
	// ルーティングごとに発生する Handler の定義と
	// レシーバー付き ServeHTTP の定義を省略できる

	// i.e. ハンドラ関数 HandleFunc は
	// 単に Handler を作成するショートカットのようなもの
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
