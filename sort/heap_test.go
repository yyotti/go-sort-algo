package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	test(t, HeapSort)
}

func BenchmarkHeapSort_Sort1K(b *testing.B) {
	bench(b, HeapSort, slice1K)
}

func BenchmarkHeapSort_Sort16K(b *testing.B) {
	bench(b, HeapSort, slice16K)
}

func BenchmarkHeapSort_AlmostSorted1K(b *testing.B) {
	bench(b, HeapSort, almostSorted1K)
}

func BenchmarkHeapSort_AlmostSorted16K(b *testing.B) {
	bench(b, HeapSort, almostSorted16K)
}
