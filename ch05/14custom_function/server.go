// リスト5.14
package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	// funcMap 型マップを生成する
	funcMap := template.FuncMap{"fdate": formatDate}

	// New() でテンプレート生成時に渡した名前と
	// ParseFile() で解析するファイル名を一致させる
	// New() でテンプレート生成後
	// Funcs() で FuncMap をテンプレートに付与し解析する
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
