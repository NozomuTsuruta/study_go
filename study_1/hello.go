// 実行可能なプログラムは全てmainを含める
package main

// 必要なライブラリをインポート
import (
	"fmt"      // Fprintfなどのフォーマット付きのI/Oをサポート
	"net/http" // httpでやり取りするために必要
)

func handler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "Hello World, %s!", request.URL.Path[1:]) // [1:] ワイルドカードのようなパス
}

// main関数が必ず一つ必要
func main() {
	http.HandleFunc("/", handler)     // ルートURLが呼び出された時、handlerが実行
	http.ListenAndServe(":8080", nil) // null
}
