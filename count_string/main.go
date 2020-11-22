package main

import (
	"fmt"
	"strings"
)

func main(){
	counts := map[string]int{}

	str := "hgdjskhgjksdrhreiowgbiverhn"
	for _,s := range strings.Split(str,"") {
		counts[s]++ // ゼロ値が0なのでそのまま使える
	}
	fmt.Println(counts)
}