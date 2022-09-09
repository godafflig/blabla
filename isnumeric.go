package piscine

func IsNumeric(s string) bool {
	alphabet := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var count int = 0

	for i = 0; i < len(s)-1; i++ {
		for j = 0; j <= len(alphabet)-1; j++ {
			if rune(== string(s[i]) {
				count++
			}
		}
	}
	if count == len(s) {
		return true
	} else {
		return false
	}
}
