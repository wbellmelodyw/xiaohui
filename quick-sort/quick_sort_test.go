package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{4, 4, 6, 5, 3, 2, 8, 1}
	quickSort(&a, 0, len(a)-1)
	fmt.Println(a)
	b := []int{4, 4, 6, 5, 3, 2, 8, 1}
	quickSortSingleSide(&b, 0, len(b)-1)
	fmt.Println(b)
	c := []int{4, 4, 6, 5, 3, 2, 8, 1}
	quickSortByStack(&c, 0, len(c)-1)
	fmt.Println(c)
}
