package main

import (
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                                                   // postの値を取得
	user, _ := data.UserByEmail(r.PostFormValue("email"))           // メールアドレスを与えられると構造体Userを返す
	if user.Password == data.Encrtpt(r.PostFormValue("password")) { //暗号化
		session := user.CreateSession() // session作成
		cookie := http.Cookie{
			Name:      "_cookie", // 任意
			Value:     session.Uuid, // ランダムに生成されたユニークなID
			HttpOnly: true, // クッキーへのアクセスをHTTP(S)のみに制限→JSなどの非HTTPのAPIは禁止
		}
		http.SetCookie(w, &cookie) // アドレスを取得
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
