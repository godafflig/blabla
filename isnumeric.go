package piscine

func IsNumeric(s string) bool {
	alphabet := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var count int = 0
	var i int
	var j int
	for i = 0; i < len(alphabet)-1; i++ {
		for j = 0; j < len(s); j++ {
			if string(alphabet[i]) == string(s[j]) {
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