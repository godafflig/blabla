package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.FirstRune("hello!"))
	z01.PrintRune(piscine.FirstRune("salut!"))
	z01.PrintRune(piscine.FirstRune("ola!"))
	z01.PrintRune('\n')
}
