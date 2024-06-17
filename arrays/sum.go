package sum

func Sum(values []int) int {
	var result int

	for _, v := range values {
		result += v
	}

	return result
}

func SumAll(values ...[]int) []int {
	var sums []int

	for _, v := range values {
		sums = append(sums, Sum(v))
	}

	return sums
}

func SumAllTails(values ...[]int) []int {
	var sums []int

	for _, v := range values {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
