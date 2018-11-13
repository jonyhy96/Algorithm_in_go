package selection

// SelectionSort is SelectionSort imply in go
func SelectionSort(args []int) []int {
	var result = make([]int, len(args))
	copy(result, args)
	for i := 0; i < len(result); i++ {
		var temp int
		pos, max := findMax(result[:len(result)-i])
		temp = result[(len(result) - i - 1)]
		result[(len(result) - i - 1)] = max
		result[pos] = temp
	}
	return result
}

func findMax(args []int) (int, int) {
	var pos int
	var max int
	for i := 0; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
			pos = i
		}
	}
	return pos, max
}
