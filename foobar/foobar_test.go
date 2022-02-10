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

type fakeIntn int

func (i fakeIntn) Intn(n int) int {
	return int(i)
}

func TestRandomFooBar(t *testing.T) {
	// src := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(src)

	want := "Foo-Bar-14-Foo"
	var fake fakeIntn = 2

	get := RandomFooBar(fake)

	if want != get {
		t.Errorf("given random returns 2 want %q but got %q\n", want, get)
	}
}
