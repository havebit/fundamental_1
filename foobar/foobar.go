package foobar

import (
	"strconv"
)

type Intner interface {
	Intn(n int) int
}

func RandomFooBar(r Intner) string {
	n := r.Intn(50) + 1
	return say(n)
}

func SayAny(i interface{}) string {
	switch v := i.(type) {
	case int:
		return say(v)
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			n = 0
		}
		return say(n)
	}

	return ""
}

func say(given int) string {
	switch {
	case given%15 == 0:
		return "FooBar"
	case given%5 == 0:
		return "Bar"
	case given%3 == 0:
		return "Foo"
	default:
		return strconv.Itoa(given)
	}
}
