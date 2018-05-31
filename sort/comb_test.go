package sort

import (
	"testing"
)

func TestCombSort(t *testing.T) {
	test(t, CombSort)
}

func BenchmarkCombSort_Sort1K(b *testing.B) {
	bench(b, CombSort, slice1K)
}

func BenchmarkCombSort_Sort16K(b *testing.B) {
	bench(b, CombSort, slice16K)
}

func BenchmarkCombSort_AlmostSorted1K(b *testing.B) {
	bench(b, CombSort, almostSorted1K)
}

func BenchmarkCombSort_AlmostSorted16K(b *testing.B) {
	bench(b, CombSort, almostSorted16K)
}
