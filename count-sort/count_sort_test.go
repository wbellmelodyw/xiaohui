package count_sort

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	sort := countSort([]int{4, 4, 6, 5, 3, 2, 8, 1, 7, 5, 6, 0, 10})
	sort2 := countSort([]int{95, 94, 91, 98, 99, 90, 99, 93, 91, 92})
	sort3 := newCountSort([]int{95, 94, 91, 98, 99, 90, 99, 93, 91, 92})
	fmt.Println(sort)
	fmt.Println(sort2)
	fmt.Println(sort3)
}
