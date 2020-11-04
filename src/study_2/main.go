package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()                                 // マルチプレクサ(→リクエストをハンドラにリダイレクト、静的なファイルの返送)を生成
	files := http.FileServer(http.Dir("/public"))             // publicを起点とするファイル
	mux.Handle("/static", http.StripPrefix("/static", files)) // リクエストURLから頭文字列を削除

	mux.HandleFunc("/", index) // ルートuRLをハンドラ関数にリダイレクト
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signup_account)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{ // サーバーを立てる
		Addr:    "0.0.0.8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
