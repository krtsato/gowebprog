// リスト4.17
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name: "flash",

		// メッセージは空白を含むため URL エンコードが必要
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "メッセージがありません。")
		}
	} else {
		// 既存のクッキーを置き換える
		rc := http.Cookie{
			Name:    "flash",         // 同名
			MaxAge:  -1,              // 1 度利用して直ちにクッキーを削除
			Expires: time.Unix(1, 0), // 過去の Unix 時間 1970-01-01 00:00:01 を指定
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
