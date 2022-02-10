package foobar

import (
	"fmt"
	"testing"
)

func TestFooBar(t *testing.T) {
	givenWant := map[int]string{
		1:  "1",
		2:  "2",
		3:  "Foo",
		6:  "Foo",
		9:  "Foo",
		12: "Foo",
		5:  "Bar",
		10: "Bar",
		20: "Bar",
		15: "FooBar",
		30: "FooBar",
	}

	for given, want := range givenWant {
		t.Run(fmt.Sprintf("give %d wants %q", given, want), func(t *testing.T) {
			get := say(given)
			if want != get {
				t.Errorf("given %d wants %q but got %q\n", given, want, get)
			}
		})
	}
}
