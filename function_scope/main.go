package main

import "fmt"

func main()  {
	fs := make([]func(),3)
	for i := range fs {
		i := i // これがないとfs[i]のfuncの中のiが決定しない
		fmt.Println("i:",i)
		fs[i]=func(){fmt.Println(i)} // func(i int) {fmt.Println(i)}(i)にしたら、iの代入がいらない
	}
	for _, f := range fs {
		f()
	}
}