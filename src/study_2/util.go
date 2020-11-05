package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie") // Cookieを取り出す
	if err == nil {                    // ログイン状態
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 空インタフェースはどの方も受け付ける、可変長引数関数
func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...)) // ParseFiles→テンプレートファイルを解析してテンプレートを作成、Mustでエラー検知
	templates.ExecuteTemplate(writter, "layout", data)
}
