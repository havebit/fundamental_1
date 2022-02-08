package foobar

import "testing"

func TestFooBarGiven1WantString1(t *testing.T) {
	given := 1
	want := "1"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}

func TestFooBarGiven2WantString2(t *testing.T) {
	given := 2
	want := "2"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}

func TestFooBarGiven3WantStringFoo(t *testing.T) {
	given := 3
	want := "Foo"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}

func TestFooBarGiven4WantString4(t *testing.T) {
	given := 4
	want := "4"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}

func TestFooBarGiven5WantStringBar(t *testing.T) {
	given := 5
	want := "Bar"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}

func TestFooBarGiven6WantStringFoo(t *testing.T) {
	given := 6
	want := "Foo"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}

func TestFooBarGiven7WantString7(t *testing.T) {
	given := 7
	want := "7"

	get := say(given)
	if want != get {
		t.Errorf("given %d wants %q but got %q\n", given, want, get)
	}
}
