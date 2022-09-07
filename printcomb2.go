package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	var col1 int = 0
	var col2 int = 1

	for col1 <= 98 {
		for col2 <= 99 {
			if col1 < 10 {
				z01.PrintRune('0')
			}
			PrintNbr(col1)
			z01.PrintRune(' ')
			if col2 < 10 {
				z01.PrintRune('0')
			}
			PrintNbr(col2)
			if col1 != 98 || col2 != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			col2 += 1
		}
		col1 += 1
		col2 = col1 + 1
	}
	z01.PrintRune('\n')
}
