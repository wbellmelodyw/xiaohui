package heap_sort

import "fmt"

//最大堆下沉
func downAdJust(a *[]int, parentIndex int, length int) {
	array := *a
	temp := array[parentIndex]
	childrenIndex := parentIndex*2 + 1
	for childrenIndex < length {
		//取最大的子节点
		if childrenIndex+1 < length && array[childrenIndex] < array[childrenIndex+1] {
			childrenIndex++
		}
		if temp >= array[childrenIndex] {
			break
		}
		array[parentIndex] = array[childrenIndex]
		parentIndex = childrenIndex
		childrenIndex = childrenIndex*2 + 1
	}
	array[parentIndex] = temp
}

func heapSort(a *[]int) {
	array := *a
	length := len(array)
	//先把无序数组构建成最大堆
	for i := (length - 2) / 2; i >= 0; i-- {
		downAdJust(a, i, length)
	}
	fmt.Println(array)
	//头和尾互换
	for i := length - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		downAdJust(a, 0, i) //避免最大被重新弄上去
	}
}
