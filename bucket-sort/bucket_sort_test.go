package bucket_sort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	fmt.Println(bucketSort([]float64{4.12, 6.421, 0.0023, 3.0, 2.123, 8.122, 4.13, 4.12, 1, 6}))
}
