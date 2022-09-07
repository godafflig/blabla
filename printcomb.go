package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var col1 int = 0
	var col2 int = 1
	var col3 int = 2

	for col1 <= 7 {
		for col2 <= 8 {
			for col3 <= 9 {
				z01.PrintRune(rune(col1 + 48))
				z01.PrintRune(rune(col2 + 48))
				z01.PrintRune(rune(col3 + 48))
				if col1 != 7 || col2 != 8 || col3 != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				col3 += 1
			}
			col2 += 1
			col3 = col2 + 1
		}
		col1 += 1
		col2 = col1 + 1
		col3 = col2 + 1
	}
	z01.PrintRune('\n')
}
