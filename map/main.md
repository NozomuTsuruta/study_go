map の作成
ゼロ値はnil

```go
変数名 := map[keyの型]valueの型{
	key:   value,
	...,
}

// 空のmap宣言
var 変数名 = map[keyの型]valueの型

変数名 := make([keyの型]valueの型,cap(任意))
```

mapの要素追加、削除

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

mapとstructsの違い
||map|structs|
|keyの型|全て同じ|stringのみ|
|valueの型|全て同じ|違う型にできる|
|keyのindexサポート|○|×|
|ポインタ|参照型|値型|

sliceはポインタも挟まっているので比較できない
mapは値を比較できる

