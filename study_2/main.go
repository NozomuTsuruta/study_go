package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	// 2つの返り値取得
	hand,remainingCards :=deal(cards, 5)

	// receiver呼び出し
	hand.print()
	remainingCards.print()
}
