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
			fmt.Println("Before:", current, "var i: ", i)
			if arr[current-1] > arr[current] {
				arr[current-1], arr[current] = arr[current], arr[current-1]
			}
			current--
			//fmt.Println("-- After:", current, "var i: ", i)
		}
	}
}
