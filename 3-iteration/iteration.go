package iteration

func Repeat(repeat string) string {
	var word string
	for i := 0; i < 5; i++ {
		word += repeat
	}

	return word
}
