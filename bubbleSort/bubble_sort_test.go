package bubbleSort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T){
	a := []int{5,8,6,3,9,2,1,7}
	fmt.Println(sort(a))
	b := []int{5,8,6,3,9,2,1,7}
	fmt.Println(tagSort(b))
}

func BenchmarkSort1(b *testing.B){
	a := []int{5,8,6,3,9,2,1,7}
	for n:=0;n<b.N;n++{
		fmt.Println(sort(a))
	}
}

func BenchmarkSort2(b *testing.B){
	s := []int{5,8,6,3,9,2,1,7}
	for n:=0;n<b.N;n++{
		fmt.Println(tagSort(s))
	}
}
func BenchmarkSort3(b *testing.B){
	s := []int{5,8,6,3,9,2,1,7}
	for n:=0;n<b.N;n++{
		fmt.Println(tagExSort(s))
	}
}
func BenchmarkSort4(b *testing.B){
	s := []int{5,8,6,3,9,2,1,7}
	for n:=0;n<b.N;n++{
		fmt.Println(cocktailSort(s))
	}
}