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
    eatBreakFast()
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