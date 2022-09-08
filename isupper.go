package piscine

func IsUpper(s string) bool {
	alphabet := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "W", "X", "Y", "Z"}
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
