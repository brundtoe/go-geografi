package main

import (
	"example.com/geografi/morestrings"
	"fmt"
)

func main() {
	fmt.Println("Hello Jackie from cmd hello")
	res := morestrings.ReverseRunes("jensen")
	fmt.Println("Jensen ", res)
	fmt.Println("Message", morestrings.Hello())
}
