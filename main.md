package ≒ スコープ

- 1 つのファイルに 1 つの package
- 1 つのディレクトリに 1 種類の package
- package main にある main 関数がエンドポイント

main は実行可能ファイルを示し必須、ないと build されない

import → 読み込む package を宣言

`go build コンパイル後のファイル名 コンパイルしたいファイル名`
→ オプション無しだと main 関数のあるファイル、カレントディレクトリ以下も自動で読み込む

`./コンパイル後のファイル名` で実行

`go run 実行したいファイル名`
→ コンパイル＋実行をやってくれるがカレントディレクトリ以下を自動で読み込まない

変数宣言(省略形は型推論)
`var 変数名 型 = 値` → `変数名 := 値`

初期化はできるが,値の代入は main 関数内でやる必要がある

関数宣言

```go
func 関数名() 戻り値の型 {
    return
}

func add(x,y int) (sum int) { // 名前付き戻り値 returnに明示しない場合につけた名前のものが返される returnするものがわかりにくい
    sum := x + y
    return
}

// 無名関数を定義後、すぐに呼び出せる
func() {
    println("hello")
}()
```

関数型

```go
num := make([]func() int,2) // int型を返す関数
num[0] = func() int {return 1}
num[1] = func() int {return 2}
```

Slice 可変配列のようなもの、同じ型しか入れられない

```go
values := string[]{"hello","world"}

or

values := make([]string,len,cap) //capを超えた場合初期値をlenを２倍にしただけのコピーが作られる 10000超えると+1/4
```

追加 → append(Slice,値)

for 文
range のあとの Slice の index と value にわけて、Slice の全ての要素を繰り返す

```go
for i := 0; i<len(values); i++ {
    fmt.Println(i)
}

// 継続条件のみ
for i <= 100 {

}

// foreach
for index, value := range values {
    fmt.Println(value)
}

// 無限ループ
for {
    break; // breakで抜け出せる
}
```

新しい型作成
ユーザー定義型　基底型と相互キャスト可能

```go
type values []string
```

receiver 関数

```go
func (引数 受け取る型) 関数名() 戻り値の型 {

}

// 呼び出す時は↓
値.関数名() ← 値は受け取る型同じ

```

Slice の値参照

```go
slice名[startIndexIncluding : endIndexNotIncluding]

slice名[: endIndexNotIncluding] // 0からendIndexNotIncludingまで

slice名[startIndexIncluding:] // startIndexIncludingから最後まで

slice名[:endIndexNotIncluding:Cap] // capacityも指定できる
```

値を 2 つ返す関数

```go
func 関数名() (戻り値1の型, 戻り値2の型) {

}
```

2 つの戻り値を受け取る

```go
変数1,変数2 := 関数名()
```

型変換

```go
型(変数など)
```

入れ替え

```go
a,b = b,a
```

structs の定義

```go
var values struct {
    value string
    index int
}{
    value: "aaa",
    index: 2,
}

// 構造体の型を定義している
type values struct {
    value string
    index int
}
```

structs の宣言

```go
val := values{"abc",3} or values{3,"abc"} // どっちがどっちかわからないので下の方がいい

val := values{value: "abc",index: 3}

or

var val values
```

型と初期値

| type   | defaultValue |
| ------ | ------------ |
| string | ""           |
| int    | 0            |
| float  | 0            |
| bool   | false        |

ポインタ
基本的に代入はコピー
代入元と同じ`値`がコピーされる
関数の引数や戻り値も

```go
val.value="bcd" ← これはコピーされたvalの値を変えている
```

変数を関数に渡す → 変数のコピーを操作している

structs の値はメモリに保存され、普通に代入するだけだと他のアドレスとしてメモリに保存されるため変わらない
|Address|Value|
|0000||
|0001|values{value: "abc"...}|
|0002||
|0003|values{value: "bcd"...} 更新した値|
|0004||

```go
valPointer := &variable // これはオリジナルのvalのアドレスを参照できる(上の図では0001)

func (pointerValues *values) updateValue() { // &で参照したvalのポインタの型を要求
    // アドレスからメモリの値を変更する
    *pointerValues.value = "bcd"
}
```

& → 値をポインタに変換、\* → ポインタを値に変換

ポインタのショートカット

```go
val.updateValue() // valの型はvaluesで、receiverの要求する型はvaluesのポインタとなっているが、自動的にvaluesのポインタに変わり、アドレスが取得できる
```

slice を作成すると slice の length,capacity,基になる配列への参照を記録する配列と slice の 2 つができる

// slice 作成 values{value: "abc",2}
[slice の length, slice の capacity, 基の配列への参照] → slice + ["abc", 2] → array

map の作成
ゼロ値は nil

```go
変数名 := map[keyの型]valueの型{
	key:   value,
	...,
}

// 空のmap宣言
var 変数名 = map[keyの型]valueの型

変数名 := make([keyの型]valueの型,cap(任意))
```

map の要素追加、削除

```go
// 作成
mapの変数名[key] = value

// 存在確認 n,okは適当な名前
n,ok := mapの変数名[key]
println(n,ok)
// 値,true or 0,false

// 削除
delete(mapの変数名,key)
```

map と structs の違い
||map|structs|
|key の型|全て同じ|string のみ|
|value の型|全て同じ|違う型にできる|
|key の index サポート|○|×|
|ポインタ|参照型|値型|

slice はポインタも挟まっているので比較できない
map は値を比較できる

関数内で値を使用しない

```go
// receiver 受け付ける型のみ記述
func (型) returnHello() string {
    return "Hello"
}
```

interface

```go
type person interface {
    eatBreakFast() string
}

type engineer struct {
    name string
}

func main() {
    nozomu := engineer{name: "nozomu"}
    print(nozomu)
}
// ↑personとして使えるように
// interfaceと同じreceiverを定義
func (e engineer) eatBreakFast() string {
    return e.name+"eat breakfast"
}

func print(p person) {
    fmt.Println(p.eatBreakFast())
}
```

型によって少し処理を変えたりできて便利

go routine

元々 1 つの routine で実行されているが、go と記述すると main routine から子 routine ができ、後の関数の内部ではて並列処理される
routine の状態は go scheduler で監視される

```go
go 関数()
```

channel
現状 main routine が実行され、子 routine ができて実行された時には main routine は処理を終了していて、結果が反映されない
go routine の間で通信を行う
channel 経由でデータを受信する時、main routine の他の操作は中断される

```go
// channel作成 値は定義したスコープ内でのみ参照可能
c := make(chan 型) // 特定の型から値を作成する組み込み関数

func print(c chan 型) {
    // 値をchannelに送信
    c<-value

    // channelの値を代入
    value<-channel

    // channelの値をprint
    fmt.Println(<-c)
}
```
