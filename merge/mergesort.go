package merge

func Sort(args []int) []int {
	if len(args) <= 1 {
		return args
	}
	var result []int
	middle := len(args) / 2
	leftArray := Sort(args[:middle])
	rightArray := Sort(args[middle:])
	merge(leftArray, rightArray, &result)
	return result
}

func merge(leftArray []int, rightArray []int, result *[]int) {
	i, j := 0, 0
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			*result = append(*result, leftArray[i])
			i++
		} else {
			*result = append(*result, rightArray[j])
			j++
		}
	}
	for ; i < len(leftArray); i++ {
		*result = append(*result, leftArray[i])
	}
	for ; j < len(rightArray); j++ {
		*result = append(*result, rightArray[j])
	}
}
