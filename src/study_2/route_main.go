package main

import (
	"net/http"
)

// クライアントに送信するHTMLの生成
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r) // 不要な情報はブランク識別子を使う
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
		// テンプレートファイルからコンテンツを取り出し、別のデータと組み合わせてHTMLを生成→テンプレートの実行
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
