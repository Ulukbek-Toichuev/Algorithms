package main

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/Ulukbek-Toychuev/Algorithms/sort"
)

func main() {

	var array1, array2 []int
	var num int = 200000

	rand.Seed(time.Now().UnixNano())
	array1 = append(array1, rand.Perm(num)...)

	beginTime1 := time.Now().UnixNano() / int64(time.Millisecond)
	s.BubbleSort(array1)
	endTime1 := time.Now().UnixNano() / int64(time.Millisecond)

	sum1 := endTime1 - beginTime1
	if sum1 > 1000 {
		sum1 = sum1 / 1000
	}
	fmt.Println(sum1)

	rand.Seed(time.Now().UnixNano())
	array2 = append(array2, rand.Perm(num)...)

	beginTime2 := time.Now().UnixNano() / int64(time.Millisecond)
	s.SelectSort(array1)
	endTime2 := time.Now().UnixNano() / int64(time.Millisecond)

	sum2 := endTime2 - beginTime2
	if sum2 > 1000 {
		sum2 = sum2 / 1000
	}
	fmt.Println(sum2)

}
