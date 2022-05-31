package main

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/Ulukbek-Toychuev/Algorithms/sort"
)

func main() {

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

	array1 := generateSlice(5)
	fmt.Println("With func generateSlice()")
	fmt.Println(array1)

	beginTime2 := time.Now().UnixNano() / int64(time.Millisecond)
	s.InsertionSort(array1)
	endTime2 := time.Now().UnixNano() / int64(time.Millisecond)

	sum2 := endTime2 - beginTime2
	if sum2 > 1000 {
		sum2 = sum2 / 1000
		fmt.Println(sum2, " second")
	} else {
		fmt.Println(sum2, "millisecond")
	}
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}
