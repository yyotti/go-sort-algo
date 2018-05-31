package sort

import (
	"testing"
)

func TestShakerSort(t *testing.T) {
	test(t, ShakerSort)
}

func BenchmarkShakerSort_Sort1K(b *testing.B) {
	bench(b, ShakerSort, slice1K)
}

func BenchmarkShakerSort_Sort16K(b *testing.B) {
	bench(b, ShakerSort, slice16K)
}

func BenchmarkShakerSort_AlmostSorted1K(b *testing.B) {
	bench(b, ShakerSort, almostSorted1K)
}

func BenchmarkShakerSort_AlmostSorted16K(b *testing.B) {
	bench(b, ShakerSort, almostSorted16K)
}
