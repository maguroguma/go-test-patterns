package animals

import (
	"testing"

	"github.com/my0k/go-test-patterns/dena-swet/foods"
)

func TestDuck05(t *testing.T) {
	t.Run("it says quack", func(t *testing.T) {
		duck := createInstance(t)

		actual := duck.Say()
		expected := "tarou says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		duck := createInstance(t)

		apple := foods.NewApple("sunfuji")

		actual := duck.Eat(apple)
		expected := "tarou ate sunfuji"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}

func createInstance(tb testing.TB) *Duck {
	tb.Helper()

	duck := NewDuck("tarou")
	// time.Sleep(5 * time.Second)
	// tb.Error("前処理で失敗しました")

	return duck
}
