package main

import (
	"fmt"
)

func main() {
	var test = []int{5, 4, 3, 2, 1}
	fmt.Println(test)
	fmt.Println(BubbleSort(test))
	fmt.Println(BubbleSort2(test))
}

// BubbleSort is a function to imply Bubblesort in golang
func BubbleSort(args []int) []int {
	fmt.Println("BubbleSort1 执行")
	var result = make([]int, len(args))
	copy(result, args)
	var isChanged bool
	var looptime int
	for i := 0; i < len(result); i++ {
		isChanged = false
		for j := i + 1; j < len(result); j++ {
			looptime++
			if result[i] > result[j] {
				changeValue(&result[i], &result[j])
				isChanged = true
				fmt.Println("交换后排序情况", result)
			}
		}
		if isChanged == false {
			fmt.Println("循环次数", looptime)
			break
		}
	}
	return result
}

// BubbleSort2 is a function to imply Bubblesort in golang
func BubbleSort2(args []int) []int {
	fmt.Println("BubbleSort2 执行")
	var result = make([]int, len(args))
	copy(result, args)
	var isChanged bool
	var looptime int
	for i := 0; i < len(result); i++ {
		isChanged = false
		for j := 0; j < len(result)-i-1; j++ {
			looptime++
			if result[j] > result[j+1] {
				changeValue(&result[j], &result[j+1])
				isChanged = true
				fmt.Println("交换后排序情况", result)
			}
		}
		if isChanged == false {
			fmt.Println("循环次数", looptime)
			break
		}
	}
	return result
}
func changeValue(valueA *int, valueB *int) {
	var temp int
	temp = *valueB
	*valueB = *valueA
	*valueA = temp
}
