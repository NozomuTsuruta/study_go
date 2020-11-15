package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_card")

	cards = newDeckFromFile("my_card")

	// 2つの返り値取得
	hand, remainingCards := deal(cards, 5)

	// receiver呼び出し
	hand.print()
	remainingCards.print()
}
