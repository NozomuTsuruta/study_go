map の作成

```go
変数名 := map[keyの型]valueの型{
	key:   value,
	...,
}

// 空のmap宣言
var 変数名 = map[keyの型]valueの型

変数名 := make([keyの型]valueの型)
```

mapの要素追加、削除

```go
// 作成
mapの変数名[key] = value

// 削除
delete(mapの変数名,key)
```

mapとstructsの違い
||map|structs|
|keyの型|全て同じ|stringのみ|
|valueの型|全て同じ|違う型にできる|
|keyのindexサポート|○|×|
|ポインタ|参照型|値型|
