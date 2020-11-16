package main

import (
	"fmt"
)

// 構造体を使わない場合はそのまま値を更新できる
func main() {
	mySlice := []string{"Hello", "World"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "The"
}
