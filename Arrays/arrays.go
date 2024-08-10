package arrays

func Sum(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func SumAll(slice ...[]int) []int {
	numberOfSlices := len(slice)

	result := make([]int, numberOfSlices)
	for i := 0; i < numberOfSlices; i++ {
		for j := 0; j < len(slice[i]); j++ {
			result[i] += slice[i][j]
		}
	}
	return result
}

func SumAllNew(slices ...[]int) []int {
	numberOfSlices := len(slices)

	result := make([]int, numberOfSlices)
	for i, slice := range slices {
		result[i] = Sum(slice)
	}
	return result
}

func SumAppend(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		result = append(result, Sum(slice))
	}
	return result
}

func SumTail(slices ...[]int) []int {
	result := make([]int, len(slices))
	for i, slice := range slices {
		if len(slice) == 0 {
			continue
		}
		result[i] = slice[len(slice)-1]
	}
	return result
}
