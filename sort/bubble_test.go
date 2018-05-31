package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	test(t, BubbleSort)
}

func BenchmarkBubbleSort_Sort1K(b *testing.B) {
	bench(b, BubbleSort, slice1K)
}

func BenchmarkBubbleSort_Sort16K(b *testing.B) {
	bench(b, BubbleSort, slice16K)
}

func BenchmarkBubbleSort_AlmostSorted1K(b *testing.B) {
	bench(b, BubbleSort, almostSorted1K)
}

func BenchmarkBubbleSort_AlmostSorted16K(b *testing.B) {
	bench(b, BubbleSort, almostSorted16K)
}
