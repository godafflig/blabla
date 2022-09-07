package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nb int) {
	var i int = 1

	if nb == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}
	if nb < 0 {
		nb = nb * -1
		z01.PrintRune('-')
	}
	for nb/i >= 10 {
		i = i * 10
	}
	for i > 0 {
		z01.PrintRune(rune((nb-(nb%i))/i + 48))
		nb = nb % i
		i = i / 10
	}
}
