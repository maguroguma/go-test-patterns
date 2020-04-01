# Goのテストパターン

---

## 基本

写経元: https://swet.dena.com/entry/2018/01/16/211035

### アサーション

```go
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
```

アサーションのメッセージも最初はこのようなシンプルな形で。

### テスト実行

```sh
go test -v -cover ./dena-swet/animals
```

Go Modulesを有効にしておくと、シンプルにパッケージ名を指定する形で実行できる。
あるいは、パッケージのディレクトリで `go test -v -cover` でもよい。

### サブテスト

テストを階層構造にすることができる。

```go
t.Run("{{ sub-test-name }}", func(t *testing.T) {
  // 通常のテスト
  ...
})
```

---

## テストにおける共通処理

写経元: https://swet.dena.com/entry/2018/01/22/120155

### BeforeAll / AfterAll

`func TestMain(m *testing.M)` を利用すると、
**すべてのテストケースの前と後に実行する `BeforeAll, AfterAll` を実現できる。**

**テストファイル内に `TestMain` が存在している場合、 `go test` は `TestMain` のみ実行する。**
`testing.M` の `Run` メソッドを呼ぶことで各テストケースが実行され、失敗か成功かに応じて
コードを返却する。

最終的に `os.Exit` に `0` が渡ればそのテストファイルは成功、それ以外の値の場合は失敗になる。

```go
func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}
```

※慣れないうちは上記の形式に従ったほうが良さそう。

### BeforeEach / AfterEach

「ある」テストケースごとに前処理や後処理を実行する方法。

**「関数による共通化」と「共通メソッドのエラー処理」を行う。**

#### 共通メソッドのエラー処理

`testing.TB` インタフェースを渡してやる。

`Fatal` をコールすることでテストを失敗にし、その場でそのテストケースを終了させてやることができる。
`Error` の場合はテストを失敗にするものの、テストケースは引き続き実行される。

```go
func createInstance(tb testing.TB) *Duck {
	tb.Helper()

	duck := NewDuck("tarou")
	// time.Sleep(5 * time.Second)
	tb.Error("前処理で失敗しました")

	return duck
}
```

`tb.Helper()` を置いておくと、失敗時の呼び出し元がわかりやすくなる。

---

## gomockを使ったテスト

写経元: https://swet.dena.com/entry/2018/01/29/141707

### モックの意義

分割したアプリケーションのそれぞれのコードをテストする場合、
モックを利用して外部（テスト対象以外）の依存を減らすことが望ましい。

`gomock` では静的なモック用のソースコードを生成する。

### モックの作成

**モックはインタフェースに対して作成する！**

以下のような `mockgen` コマンドで作成する。

```
$ mockgen -source=foods/food.go --destination foods/mock_foods/mock_foods.go
```

### モックを使ったテスト

以下はモックの作成部分。

```go
ctrl := gomock.NewController(t)
food := mock_foods.NewMockFood(ctrl)

food.EXPECT().Name().Return("kougyoku")
```

テストコード全体は、以下のような形になる。

```go
func TestDuck_Eat_02(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    food := mock_foods.NewMockFood(ctrl)
    food.EXPECT().Name().Return("kougyoku")

    duck := animals.NewDuck("tarou")
    actual := duck.Eat(food)
    expected := "tarou ate kougyoku"
    if actual != expected {
        t.Errorf("got: %s\nwont: %s", actual, expected)
    }
}
```

> `Food` はインタフェースなので本来であればその実装に依存しますが、
> **モックを作成して特定の文字列を返却するように指定している** のでテスト実行時の依存を減らすことができます。

※モックを擬似的なもので指定している雰囲気を掴むのが大事。

引数付きのものも、引数ごとに色々と定義できるっぽい。

```go
type Hoge interface {
    Foo(foo string) string
}

hogeMock.EXPECT().Foo(gomock.Any()).Return("bar")
```

