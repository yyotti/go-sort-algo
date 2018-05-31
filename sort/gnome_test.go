package sort

import (
	"testing"
)

func TestGnomeSort(t *testing.T) {
	test(t, GnomeSort)
}

func BenchmarkGnomeSort_Sort1K(b *testing.B) {
	bench(b, GnomeSort, slice1K)
}

func BenchmarkGnomeSort_Sort16K(b *testing.B) {
	bench(b, GnomeSort, slice16K)
}

func BenchmarkGnomeSort_AlmostSorted1K(b *testing.B) {
	bench(b, GnomeSort, almostSorted1K)
}

func BenchmarkGnomeSort_AlmostSorted16K(b *testing.B) {
	bench(b, GnomeSort, almostSorted16K)
}
