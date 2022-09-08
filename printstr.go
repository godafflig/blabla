package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	var i int
	for i = 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}
