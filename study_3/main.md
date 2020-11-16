構造体の定義

```go
type values struct {
    value string
    index int
}
```

構造体の宣言

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
val.value="bcd" ← これはvalをコピーしてそれの値を変えている
```

構造体の値はメモリに保存され、普通に代入するだけだと他のアドレスとしてメモリに保存されるため変わらない
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
