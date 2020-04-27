package heap_sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	a := &[]int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	heapSort(a)
	fmt.Println(a)
}
