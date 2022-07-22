package goutils

func Getmax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func NumberToWord(a int) string {
	var numbers := []string {"bir", "ikki", "uch", "tort", "besh", "olti", "yetti", "sakkiz", "tuqqiz"}
	
	if a > 0 && a < 10 {
		return numbers[a]
	} else {
		return "Bu son 1 va 10 oralig'ida emas!"
	}
}
