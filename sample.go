package main

import (
	"log"

	"github.com/go-collection/collection"
)

func main() {
	str := collection.ListOf(1, 2, 3).Map(
		func(e interface{}) interface{} {
			num := e.(int)
			return num * 10
		},
		func(e interface{}) interface{} {
			num := e.(int)
			return num + 900
		},
	).ToString()

	log.Println(str)
}
