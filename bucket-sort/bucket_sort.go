package bucket_sort

import (
	"sort"
)

//桶排序
//区间跨度 =（max-min)/(length-1)
//元素 = （now-min）/跨度区间
func bucketSort(array []float64) []float64 {
	bucketNum := len(array)
	min := array[0]
	max := array[0]
	//找出最大最小元素
	for i := 1; i < bucketNum; i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}
	//初始化桶
	area := (max - min) / float64(bucketNum-1)

	buckets := make([][]float64, bucketNum) //创建对应数组个数的桶
	for i := 0; i < bucketNum; i++ {
		index := int((array[i] - min) / area)
		buckets[index] = append(buckets[index], array[i])
	}
	//对每个桶内部进行排序
	for _, bucket := range buckets {
		sort.Float64s(bucket)
	}
	//分别输出所有桶

	sortArray := make([]float64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			for _, item := range bucket {
				sortArray = append(sortArray, item)
			}
		}
	}
	return sortArray
}
