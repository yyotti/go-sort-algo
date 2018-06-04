package sort

import (
	"testing"
)

func TestParallelMergeSort(t *testing.T) {
	test(t, ParallelMergeSort)
}

func BenchmarkParallelMergeSort_Sort1K(b *testing.B) {
	bench(b, ParallelMergeSort, slice1K)
}

func BenchmarkParallelMergeSort_Sort16K(b *testing.B) {
	bench(b, ParallelMergeSort, slice16K)
}

func BenchmarkParallelMergeSort_AlmostSorted1K(b *testing.B) {
	bench(b, ParallelMergeSort, almostSorted1K)
}

func BenchmarkParallelMergeSort_AlmostSorted16K(b *testing.B) {
	bench(b, ParallelMergeSort, almostSorted16K)
}
