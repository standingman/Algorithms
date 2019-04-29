package sort
/**
	推排序
 */
func BuildHeap(arr []int, length int) {
	if length == 1 {
		return
	}
	maxbrach := length/2 - 1
	for i := maxbrach; i >= 0; i-- {
		lChild := i*2 + 1
		rChild := lChild + 1
		tmpmax := i
		if arr[lChild] > arr[tmpmax] {
			tmpmax = lChild
		}
		if rChild <= length-1 && arr[rChild] > arr[tmpmax] {
			tmpmax = rChild
		}
		if tmpmax != i {
			arr[i], arr[tmpmax] = arr[tmpmax], arr[i]
		}
	}
}
func HeapSort(arr []int){
	length := len(arr)
	for i := 0; i < length; i++ {
		lastMessLen:=length-i
		BuildHeap(arr,lastMessLen)
		if i<length{
			arr[0],arr[lastMessLen-1]=arr[lastMessLen-1],arr[0]
		}
	}
}
