package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	test(t, QuickSort)
}

func BenchmarkQuickSort_Sort1K(b *testing.B) {
	bench(b, QuickSort, slice1K)
}

func BenchmarkQuickSort_Sort16K(b *testing.B) {
	bench(b, QuickSort, slice16K)
}

func BenchmarkQuickSort_AlmostSorted1K(b *testing.B) {
	bench(b, QuickSort, almostSorted1K)
}

func BenchmarkQuickSort_AlmostSorted16K(b *testing.B) {
	bench(b, QuickSort, almostSorted16K)
}
