// リスト5.28
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template

	// Layout ファイルで共通部分を描画した上で
	// 条件分岐するコンポーネントは別のテンプレートとして切り分ける
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
