package sort
/**
冒泡排序
 */
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
