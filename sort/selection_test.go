package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	test(t, SelectionSort)
}

func BenchmarkSelectionSort_Sort1K(b *testing.B) {
	bench(b, SelectionSort, slice1K)
}

func BenchmarkSelectionSort_Sort16K(b *testing.B) {
	bench(b, SelectionSort, slice16K)
}

func BenchmarkSelectionSort_AlmostSorted1K(b *testing.B) {
	bench(b, SelectionSort, almostSorted1K)
}

func BenchmarkSelectionSort_AlmostSorted16K(b *testing.B) {
	bench(b, SelectionSort, almostSorted16K)
}
