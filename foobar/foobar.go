package foobar

import "strconv"

func say(given int) string {

	if given == 5 {
		return "Bar"
	}

	if given%3 == 0 {
		return "Foo"
	}

	return strconv.Itoa(given)
}
