package insertion

// Sort is a function to imply InsertionSort in golang
func Sort(args []int) []int {
	var result = make([]int, len(args))
	var looptime int
	copy(result, args)
	var temp int
	for i := 1; i < len(result); i++ {
		temp = result[i]
		for i > 0 && result[i-1] > temp {
			looptime++
			result[i] = result[i-1]
			i--
		}
		result[i] = temp
	}
	return result
}
