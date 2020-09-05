// リスト6.18
/*
 1. go get github.com/jinzhu/gorm
 2. go run store.go
*/
package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"` // Gorm に対して not null を指定
	Comments  []Comment
	CreatedAt time.Time
}

// Warning: Comment 構造体に Post フィールドがない
// Gorm は PostId フィールドを自動的に外部キーだと設定する
type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"` // Gorm に対して not null を指定
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

// connect to the Db
func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}

	// AutoMigrate は可変長引数
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post) // {0 Hello World! Sau Sheong [] 0001-01-01 00:00:00 +0000 UTC}

	// 構造体 Post のメソッド Create() を作る必要はない
	Db.Create(&post)
	fmt.Println(post) // {1 Hello World! Sau Sheong [] 2015-04-13 11:38:50.91815604 +0800 SGT}

	// Add a comment
	comment := Comment{Content: "いい投稿だね！", Author: "Joe"}

	// Post モデルを選択し, Comment を関連付けて, Comment を追加する
	// PostId に手動でアクセスしていない
	Db.Model(&post).Association("Comments").Append(comment)

	// Get comments from a post
	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)

	var comments []Comment

	// readPost に関連する comments を Related() で取得する
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0]) // {1 Good post! Joe 1 2015-04-13 11:38:50.920377 +0800 SGT}
}
