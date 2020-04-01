package animals

import (
	"os"
	"testing"

	"github.com/my0k/go-test-patterns/dena-swet/foods"
)

func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestDuckSay03(t *testing.T) {
	duck := NewDuck("tarou")
	actual := duck.Say()
	expected := "tarou says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestDuckEat(t *testing.T) {
	duck := NewDuck("tarou")
	apple := foods.NewApple("sunfuji")

	actual := duck.Eat(apple)
	expected := "tarou ate sunfuji"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
