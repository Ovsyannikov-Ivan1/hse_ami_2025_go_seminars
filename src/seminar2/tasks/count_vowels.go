package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(s string) int {
	ans := 0
	for _, ch := range s {
		if ch == 'e' || ch == 'y' || ch == 'u' || ch == 'i' || ch == 'o' || ch == 'a' || ch == 'E' || ch == 'Y' || ch == 'U' || ch == 'J' || ch == 'O' || ch == 'A' {
			ans++
		}
	}
	return ans
}
