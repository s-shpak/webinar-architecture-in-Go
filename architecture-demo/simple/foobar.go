package main

import "strconv"

func GetFooBar(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FooBar"
	}
	if n%3 == 0 {
		return "Foo"
	}
	if n%5 == 0 {
		return "Bar"
	}
	return strconv.Itoa(n)
}
