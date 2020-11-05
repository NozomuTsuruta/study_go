package main

import (
	"html/template"
	"net/http"
)

// クライアントに送信するHTMLの生成
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r) // 不要な情報はブランク識別子を使う
		publicTmplFiles := []string{"templates/layout.html", "templates/public.navbar.html", "templates/index.html"}
		privateTmplFiles := []string{"templates/layout.html", "templates/public.navbar.html", "templates/index.html"}

		var templates *template.Template
		if err != nil {
			// ParseFiles→テンプレートファイルを解析してテンプレートを作成、Mustでエラー検知
			templates = template.Must(template.ParseFiles(publicTmplFiles...))
		} else {
			templates = template.Must(template.ParseFiles(privateTmplFiles...))
		}
		// テンプレートファイルからコンテンツを取り出し、別のデータと組み合わせてHTMLを生成→テンプレートの実行
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
