// リスト4.16
package main

import (
	"fmt"
	"net/http"
)

/*
Go 言語における Cookie の構造体
type Cookie struct {
	Name string
	Value string
	Plan string
	Domain string
	Expires time.Time
	RawExpires string
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}

Expires フィールドが設定されていない場合
・セッションクッキー (一時的)
・ブラウザが閉じられるとブラウザから削除される

Expires フィールドが設定されている場合
・期限切れ・削除されない限り持続する

有効期限の指定方法
・Expires フィールドを使う : いつ期限切れになるか
・MaxAge フィールドを使う : ブラウザ内で生成されてから何秒間有効か
		・MaxAge=0 means no 'Max-Age' attribute specified.
		・MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
		・MaxAge>0 means Max-Age attribute present and given in seconds

Expires と MaxAge が混在している理由
・ブラウザ間のクッキー実装の不一致のため
・HTTP 1.1 では MaxAge が優先され Expires は非推奨
・しかし Expires はほぼ全てのブラウザが未だにサポートしている
・MaxAge は IE 6, 7 で未だにサポートされていない

Expires と MaxAge どちらを使うか
・Expires フィールドだけ使う
・あらゆるブラウザをサポートするため両方使う
*/

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// 複数のクッキーを取得したい場合 r.Cookies() でスライスを得る
	// 指定したクッキーが存在しない場合エラー
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
