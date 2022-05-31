package sort

//Сортировка выбором
func SelectSort(array []int) {
	len := len(array)

	for i := 0; i < len; i++ {

		minElement := i

		for j := i + 1; j < len; j++ {
			if array[j] < array[minElement] {
				minElement = j
			}
		}
		array[i], array[minElement] = array[minElement], array[i]
	}
}

//Пузырьковая сортировка
func BubbleSort(array []int) {
	len := len(array)
	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {

			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
