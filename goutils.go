package goutils

// Function to get max of two integers
func Getmax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

<<<<<<< HEAD
// Function to convert digits into uzbek words
=======
//  this is another commit

>>>>>>> 8a32f792edd164ee7db6c7d152a762d28fde6538
func NumberToWord(a int) string {
	var numbers = []string{"nol", "bir", "ikki", "uch", "tort", "besh", "olti", "yetti", "sakkiz", "tuqqiz"}

	if a >= 0 && a < 10 {
		return numbers[a]
	} else {
		return "Bu son 1 va 10 oralig'ida emas!"
	}
}
