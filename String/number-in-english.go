package main

func NumberInEnglish(num int) string {
	// inbuilt method
	// str := num2words.Convert(num)
	// return str

	if num == 0 {
		return "Zero"
	}

	values := []int{1000000000, 1000000, 1000, 100, 90, 80, 70,
		60, 50, 40, 30, 20, 19, 18, 17, 16, 15, 14,
		13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	words := []string{"Billion", "Million", "Thousand", "Hundred",
		"Ninety", "Eighty", "Seventy", "Sixty", "Fifty",
		"Forty", "Thirty", "Twenty", "Nineteen",
		"Eighteen", "Seventeen", "Sixteen", "Fifteen",
		"Fourteen", "Thirteen", "Twelve", "Eleven",
		"Ten", "Nine", "Eight", "Seven", "Six", "Five",
		"Four", "Three", "Two", "One"}

	return numberInEnglishRec(num, values, words)
}

func numberInEnglishRec(num int, values []int, words []string) string {
	res := ""

	for i := 0; i < len(values); i++ {
		if num >= values[i] {
			if num >= 100 {
				res += numberInEnglishRec(num/values[i], values, words) + " "
			}
			res += words[i]

			if num%values[i] > 0 {
				res += " " + numberInEnglishRec(num%values[i], values, words)
			}
			return res
		}
	}
	return res
}
