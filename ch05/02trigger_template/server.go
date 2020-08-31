// リスト5.2
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// HTML のテンプレートファイルを解析して構造体を生成
	t, _ := template.ParseFiles("tmpl.html")

	// ResponseWriter と何らかのデータを引数として解析済みのテンプレートを実行
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
