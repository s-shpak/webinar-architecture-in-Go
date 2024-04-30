package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := GetConfiguration()
	for i := 1; i <= cfg.Limit; i++ {
		foobar := GetFooBar(i)
		fmt.Printf("%s ", foobar)
	}
	fmt.Println()
	return nil
}
