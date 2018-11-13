package bubble

// BubbleSort is a function to imply Bubblesort in golang
func BubbleSort(args []int) []int {
	var result = make([]int, len(args))
	copy(result, args)
	var isChanged bool
	for i := 0; i < len(result); i++ {
		isChanged = false
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				changeValue(&result[i], &result[j])
				isChanged = true
			}
		}
		if i != 0 && isChanged == false {
			break
		}
	}
	return result
}

// 交换数据
func changeValue(valueA *int, valueB *int) {
	var temp int
	temp = *valueB
	*valueB = *valueA
	*valueA = temp
}
