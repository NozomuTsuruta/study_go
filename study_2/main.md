変数宣言(省略形は型推論)
`var 変数名 型 = 値` → `変数名 := 値`

初期化はできるが,値の代入は main 関数内でやる必要がある

関数宣言

```go
func 関数名() 戻り値の型 {
    return
}
```

Slice 可変配列のようなもの、同じ型しか入れられない

例)`values := string[]{"hello","world"}`

追加 → append(Slice,値)

for 文
range のあとの Slice の index と value にわけて、Slice の全ての要素を繰り返す

```go
for index, value := range values {
    fmt.Plintln(value)
}
```

新しい型作成

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
