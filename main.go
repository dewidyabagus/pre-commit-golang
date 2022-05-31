package main

import (
	"fmt"
)

func main() {
	fmt.Println(Reverse("123456789"))
}

func Reverse(text string) string {
	btext := []byte(text)

	for i, j := 0, len(btext)-1; i < len(btext)/2; i, j = i+1, j-1 {
		btext[i], btext[j] = btext[j], btext[i]
	}

	return string(btext)
}
