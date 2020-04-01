package animals

import (
	"testing"

	"github.com/my0k/go-test-patterns/dena-swet/foods"
)

func TestDuckName(t *testing.T) {
	duck := &Duck{"tarou"}
	actual := duck.name
	expected := "tarou"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestDuck(t *testing.T) {
	duck := NewDuck("tarou")

	t.Run("it says quack", func(t *testing.T) {
		actual := duck.Say()
		expected := "tarou says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		apple := foods.NewApple("sunfuji")

		actual := duck.Eat(apple)
		expected := "tarou ate sunfuji"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}
