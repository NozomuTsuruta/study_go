structs の定義

```go
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

```go
val.value="bcd" ← これはコピーされたvalの値を変えている
```

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
