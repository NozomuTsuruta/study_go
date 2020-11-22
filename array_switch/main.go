package main

import "fmt"

func main() {
	const x int = 100                                               // 型のある定数
	const n = 1_000_000_000_000_000_000 / 1_000_000_000_000_000_000 // 型のない定数
	var m = n                                                       // int型

	// 何もしないとbreak 跨ぐ際にはfallthrough
	switch m {
	case 1, 2:
		fmt.Println("m is 1 or 2") // 1 or 2
	default:
		fmt.Println("default")
	}

	// caseに式が使える
	switch {
	case m == 1:
		fmt.Println("m is 1")
	}

LOOP: // label指定で無限ループを抜け出せる
	for {
		switch {
		case m%2 == 1:
			break LOOP
		}
	}

	var ns [5]int // 要素数固定 型と要素数がセット
	println(ns)
	ms := [3]int{3, 2, 1}
	println(ms)
}
