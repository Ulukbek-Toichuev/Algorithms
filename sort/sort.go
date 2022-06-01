package sort

import "fmt"

//Сортировка выбором
func SelectSort(array []int) {
	fmt.Println("SELECT SORT")

	for i := 0; i < len(array); i++ {

		minElement := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minElement] {
				minElement = j
			}
		}
		array[i], array[minElement] = array[minElement], array[i]
	}
}

//Пузырьковая сортировка
func BubbleSort(array []int) {
	fmt.Println("BUBBLE SORT")
	for i := 0; i < len(array); i++ {

		//fmt.Println("Index ", i, "and element", array[i])
		for j := i; j < len(array); j++ {
			//fmt.Println(" -- Index ", j, "and element", array[j])
			if array[i] > array[j] {

				array[i], array[j] = array[j], array[i]

			}
		}

	}
}

func InsertionSort(arr []int) {
	fmt.Println("INSERTION SORT")
	for i := 1; i < len(arr); i++ {
		current := i
		for current > 0 {
			//fmt.Println("Before:", current, "var i: ", i)
			if arr[current-1] > arr[current] {
				arr[current-1], arr[current] = arr[current], arr[current-1]
			}
			current--
			//fmt.Println("-- After:", current, "var i: ", i)
		}
	}
}

func MergeSort(items []int) []int {
	//fmt.Println("MERGE SORT")
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
