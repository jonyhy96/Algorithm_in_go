package main

import (
	"Algorithm_in_go/bubble"
	"Algorithm_in_go/insertion"
	"Algorithm_in_go/merge"
	"Algorithm_in_go/quick"
	"Algorithm_in_go/selection"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var test = []int{4, 36, 98, 69, 83}
	// test = initArray()
	fmt.Println("代排序数据", test)
	fmt.Println("冒泡排序结果", bubble.BubbleSort(test))
	fmt.Println("选择排序结果", selection.SelectionSort(test))
	fmt.Println("插入排序结果", insertion.Sort(test))
	fmt.Println("快速排序结果", quick.Sort(test))
	fmt.Println("归并排序结果", merge.Sort(test))
	fmt.Println(check(bubble.BubbleSort(test)))
	fmt.Println(check(selection.SelectionSort(test)))
	fmt.Println(check(insertion.Sort(test)))
	fmt.Println(check(quick.Sort(test)))
	fmt.Println(check(merge.Sort(test)))
}
func initArray() []int {
	randSource := rand.NewSource(time.Now().Unix())
	randmum := rand.New(randSource)
	var args []int
	for i := 0; i < 5; i++ {
		args = append(args, randmum.Intn(100))
	}
	fmt.Println("生成的切片是", args)
	return args
}
func check(arg []int) bool {
	var right = true
	for i := 0; i < len(arg)-1; i++ {
		if arg[i] > arg[i+1] {
			right = false
		}
	}
	return right
}
