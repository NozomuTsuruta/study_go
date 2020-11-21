go routine

元々 1 つの routine で実行されているが、go と記述すると main routine から子 routine ができ、後の関数の内部ではて並列処理される
routine の状態は go scheduler で監視される

```go
go 関数()
```

channel
現状 main routine が実行され、子 routine ができて実行された時には main routine は処理を終了していて、結果が反映されない
go routine の間で通信を行う
channel経由でデータを受信する時、main routineの他の操作は中断される

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