package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	test(t, InsertionSort)
}

func BenchmarkInsertionSort_Sort1K(b *testing.B) {
	bench(b, InsertionSort, slice1K)
}

func BenchmarkInsertionSort_Sort16K(b *testing.B) {
	bench(b, InsertionSort, slice16K)
}

func BenchmarkInsertionSort_AlmostSorted1K(b *testing.B) {
	bench(b, InsertionSort, almostSorted1K)
}

func BenchmarkInsertionSort_AlmostSorted16K(b *testing.B) {
	bench(b, InsertionSort, almostSorted16K)
}
