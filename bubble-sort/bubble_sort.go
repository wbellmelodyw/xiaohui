package bubbleSort

//冒泡 从小到大
func sort(array []int) []int{
	for i:=0;i<len(array)-1;i++{//回合数
		for j:=0;j<len(array)-1-i;j++{//位置
			if array[j]>array[j+1]{
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
	}
	return array
}
//标识位优化版
func tagSort(array []int) []int{
	for i:=0;i<len(array)-1;i++{
		isSort := true
		for j:=0;j<len(array)-1-i;j++{
			if array[j]>array[j+1]{
				isSort = false
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
		if isSort{//表示不需要排序
			break
		}
	}
	return array
}
//利用回合数>=有序区数
//永远让j小于有序区
func tagExSort(array []int) []int{
	exchangeIndex := len(array)-1
	for i:=0;i<len(array)-1;i++{
		isSort := true
		for j:=0;j<exchangeIndex;j++{
			if array[j]>array[j+1]{
				isSort = false
				exchangeIndex = j
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
		if isSort{//表示不需要排序
			break
		}
	}
	return array
}

func cocktailSort(array []int) []int{
	for i:=0;i<len(array)/2;i++{
		isSort := true
		//正向
		for j:=i;j<len(array)-1-i;j++{
			if array[j]>array[j+1]{
				//发生交换
				isSort = false
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
		if isSort{//表示不需要排序
			break
		}
		isSort = true
		//逆向
		for j:=len(array)-1-i;j>i;j--{
			if array[j]<array[j-1]{
				//发生交换
				isSort = false
				array[j],array[j-1] = array[j-1],array[j]
			}
		}
		if isSort{//表示不需要排序
			break
		}
	}
	return array
}