package main

func ReverseVowels(str string) string {
	char := []rune(str)
	for i := 0; i < len(char); i++ {
		for j := i + 1; j < len(char); j++ {
			if char[i] == 'a' || char[i] == 'e' || char[i] == 'i' || char[i] == 'o' || char[i] == 'u' {
				if char[j] == 'a' || char[j] == 'e' || char[j] == 'i' || char[j] == 'o' || char[j] == 'u' {
					char[i], char[j] = char[j], char[i]
				}
			}
		}
	}
	return string(char)
}
