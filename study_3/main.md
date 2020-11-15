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
