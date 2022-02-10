package foobar

import "strconv"

func SayAny(i interface{}) string {
	if v, ok := i.(int); ok {
		return say(v)
	}
	if v, ok := i.(string); ok {
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
