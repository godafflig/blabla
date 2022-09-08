package piscine

func LastRune(s string) rune {
	var a int

	a = len(s) - 1
	return rune(s[a])
}
