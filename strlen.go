package piscine

func StrLen(s string) int {
    t1 := [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '&', 'é', '-', 'è', '_', 'ç', 'à', '!', ' '}
    var i int
    var a int = 0
    var j int
    for i = 0; i < len(s); i++ {
        for j = 0; j < len(t1); j++ {
            if rune(s[i]) == t1[j] || rune(s[i]) == 'Ã' && t1[j] == 'é' {
                a = a + 1
            }
        }
    }
    return a
}
