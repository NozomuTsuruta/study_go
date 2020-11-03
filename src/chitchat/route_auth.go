package main

import (
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // postの値を取得
	user, _ := data.UserByEmail(r.PostFormValue("email")) // メールアドレスを与えられると構造体Userを返す
	if user.Password == data.Encrtpt(r.PostFormValue)
}
