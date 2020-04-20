package heap

import (
	"fmt"
	"testing"
)

func TestUpAdJust(t *testing.T) {
	/**
	    1
	   3 2
	 6 5 7 8
	9 10 0
	*/
	a := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	upAdjust(&a)
	fmt.Println(a)
	/**
	    7
	   1 3
	 10 5 2 8
	9 6
	*/
	b := []int{7, 1, 3, 10, 5, 2, 8, 9, 6}
	buildHeap(&b)
	/**
	    1
	   5 2
	 6 7 3 8
	9 10
	*/
	fmt.Println(b)
}
