package iterations

func Repeat(word string, repeatTime int) string {
	var repeatedString string
	for i := 0; i < repeatTime; i++ {
		repeatedString += word
	}
	return repeatedString
}
