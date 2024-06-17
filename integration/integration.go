package integration

func Repeat(char string, repeatCount int) string {
	var repeated string

	for i := 1; i <= repeatCount; i++ {
		repeated += char
	}
	return repeated
}
