package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(nb int) {
	var i int = 1

	if nb < 0 {
		nb = nb * -1
		z01.PrintRune('-')
	}
	for i < nb {
		i = i * 10
	}
	i = i / 10
	for i > 0 {
		z01.PrintRune(rune((nb-(nb%i))/i + 48))
		nb = nb % i
		i = i / 10
	}
	z01.PrintRune('\n')
}
