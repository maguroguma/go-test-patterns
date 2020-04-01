package animals

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/my0k/go-test-patterns/dena-swet/foods/mock_foods"
)

func TestDuckEat02(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	food := mock_foods.NewMockFood(ctrl)
	food.EXPECT().Name().Return("kougyoku")

	duck := NewDuck("tarou")
	actual := duck.Eat(food)
	expected := "tarou ate kougyoku"
	if actual != expected {
		t.Errorf("got: %s\nwant: %s", actual, expected)
	}
}

func TestDuckEatHoge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("any arguments", func(t *testing.T) {
		hoge := mock_foods.NewMockHoge(ctrl)
		hoge.EXPECT().Foo(gomock.Any()).Return("bar")

		actual := "bar"
		expected := hoge.Foo("abababa")
		if actual != expected {
			t.Errorf("got: %s\nwant: %s", actual, expected)
		}
	})

	t.Run("for each argument", func(t *testing.T) {
		hoge := mock_foods.NewMockHoge(ctrl)
		hoge.EXPECT().Foo("hello").Return("world")
		hoge.EXPECT().Foo("I'm").Return("fine")

		actual := "world"
		expected := hoge.Foo("hello")
		if actual != expected {
			t.Errorf("got: %s\nwant: %s", actual, expected)
		}

		actual = "fine"
		expected = hoge.Foo("I'm")
		if actual != expected {
			t.Errorf("got: %s\nwant: %s", actual, expected)
		}
	})
}
