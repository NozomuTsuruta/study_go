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