package main

import (
	"fmt"

	"github.com/felipemantoan/valid-anagram/anagram"
)

func main() {
	fmt.Println(anagram.IsValid("anagram","nagaram"))
	fmt.Println(anagram.IsValid("rat","car"))
}