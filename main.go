package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println()

	/*var array1, array2 []int
	var num int = 20
	rand.Seed(time.Now().UnixNano())
	array1 = append(array1, rand.Perm(num)...)

	fmt.Println("array 1: ", array1)

	beginTime1 := time.Now().UnixNano() / int64(time.Millisecond)
	s.BubbleSort(array1)
	endTime1 := time.Now().UnixNano() / int64(time.Millisecond)

	sum1 := endTime1 - beginTime1
	if sum1 > 1000 {
		sum1 = sum1 / 1000
	}
	fmt.Println(sum1, array1)

	rand.Seed(time.Now().UnixNano())
	array2 = append(array2, rand.Perm(num)...)

	fmt.Println("array 2: ", array2)

	beginTime2 := time.Now().UnixNano() / int64(time.Millisecond)
	s.SelectSort(array2)
	endTime2 := time.Now().UnixNano() / int64(time.Millisecond)

	sum2 := endTime2 - beginTime2
	if sum2 > 1000 {
		sum2 = sum2 / 1000
	}
	fmt.Println(array2)

	var array1 []int
	var num int = 30000
	rand.Seed(time.Now().UnixNano())
	array1 = append(array1, rand.Perm(num)...)*/

	//fmt.Println("With func generateSlice()")

	//array1 := []int{3, 5, 7, 1, 8, 4, 2, 6, 4, 7, 9, 0}
	var c int
	fmt.Print("Enter element count: ")
	fmt.Scanln(&c)
	beginTime2 := time.Now().UnixNano() / int64(time.Millisecond)
	array1 := generateSlice(c)
	//array1 = s.MergeSort(array1)

	sort.Ints(array1)
	endTime2 := time.Now().UnixNano() / int64(time.Millisecond)

	sum2 := endTime2 - beginTime2
	if sum2 > 1000 {
		sum2 = sum2 / 1000
		fmt.Println("Sorted time:", sum2, " second, elements:", c)
	} else {
		fmt.Println("Sorted time:", sum2, "millisecond, elements:", c)
	}

	fmt.Println()
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}
