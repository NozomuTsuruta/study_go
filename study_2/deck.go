package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// 順番にカードを追加して、deckを返す
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver printするだけ
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver deckをstringで出力
func (d deck) toString() string {
	// join   型変換
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// bs = byte[]
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// エラー表示
		fmt.Println("Error:", err)

		// プログラムの終了
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// int64のユニークな値をseedとして渡す nano seconds time
	source := rand.NewSource(time.Now().UnixNano())

	// sourceによって、毎回違う乱数を作成するように
	r := rand.New(source)

	for i := range d {
		// 乱数作成 *Randならreceiverみたいな感じで呼べる
		newPosition := r.Intn(len(d) - 1)

		// 入れ替え
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
