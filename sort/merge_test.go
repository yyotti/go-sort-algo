package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	test(t, MergeSort)
}

func BenchmarkMergeSort_Sort1K(b *testing.B) {
	bench(b, MergeSort, slice1K)
}

func BenchmarkMergeSort_Sort16K(b *testing.B) {
	bench(b, MergeSort, slice16K)
}

func BenchmarkMergeSort_AlmostSorted1K(b *testing.B) {
	bench(b, MergeSort, almostSorted1K)
}

func BenchmarkMergeSort_AlmostSorted16K(b *testing.B) {
	bench(b, MergeSort, almostSorted16K)
}
