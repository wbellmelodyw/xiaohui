package quick_sort

//快速排序的平均时间复杂度是O(nlogn)
//但最坏情况下的时 间复杂度是O(n 2 ) 。
//1. 双边循环法。
func partition(a *[]int, startIndex, endIndex int) int {
	array := *a
	//第一个当作基准元素
	pivot := array[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		for left < right && array[right] > pivot {
			right--
		}
		for left < right && array[left] <= pivot {
			left++
		}
		//left和right交换
		if left < right { //加上少一次排序
			array[left], array[right] = array[right], array[left]
		}
	}
	array[startIndex], array[left] = array[left], array[startIndex]
	return left
}
func quickSort(a *[]int, startIndex, endIndex int) {
	//递归结束条件 一般s==e
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(a, startIndex, endIndex)
	quickSort(a, startIndex, pivotIndex-1)
	quickSort(a, pivotIndex+1, endIndex)
}

func partitionSingleSide(a *[]int, startIndex, endIndex int) int {
	array := *a
	mark := startIndex
	pivot := array[startIndex]
	for i := startIndex + 1; i <= endIndex; i++ {
		if array[i] < pivot {
			mark++
			array[mark], array[i] = array[i], array[mark]
		}
	}

	array[mark], array[startIndex] = array[startIndex], array[mark]
	return mark
}

//1. 单边循环法。
func quickSortSingleSide(a *[]int, startIndex, endIndex int) {
	//递归结束条件 一般s==e
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partitionSingleSide(a, startIndex, endIndex)
	quickSortSingleSide(a, startIndex, pivotIndex-1)
	quickSortSingleSide(a, pivotIndex+1, endIndex)
}
