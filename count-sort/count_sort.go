package count_sort

//计数排序

func countSort(a []int) []int {
	length := len(a)
	if length == 0 {
		return []int{}
	}
	min := a[0]
	max := a[0]
	//找出最大最小元素
	for i := 1; i < length; i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	//把0算上
	count := max - min + 1
	countArray := make([]int, count)
	for i := 0; i < length; i++ {
		countArray[a[i]-min]++
	}
	newArray := make([]int, 0)
	for i := 0; i < len(countArray); i++ {
		for j := 0; j < countArray[i]; j++ {
			newArray = append(newArray, i+min)
		}
	}
	return newArray
}

//优化后的计数排序
func newCountSort(a []int) []int {

	min, max := a[0], a[0]
	length := len(a)
	for i := 1; i < length; i++ {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
	}
	count := max - min + 1
	countArray := make([]int, count)
	for i := 0; i < length; i++ {
		countArray[a[i]-min]++
	}
	//标记加数字
	for i := 1; i < length; i++ {
		countArray[i] = countArray[i] + countArray[i-1]
	}

	newArray := make([]int, count)
	//倒序排序
	for i := length - 1; i >= 0; i-- {
		//countArray[a[i]-min]代表前面有几个(包括自己)
		countIndex := countArray[a[i]-min] - 1
		newArray[countIndex] = a[i]
		countArray[a[i]-min]--
	}
	return newArray
}
