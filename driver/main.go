package main

import (
	"fmt"

	"github.com/DanielSBrown/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodby", 15))
	fmt.Println(leftpad.FormatRune("Helasdfasdflo", 15, '😘'))
}
