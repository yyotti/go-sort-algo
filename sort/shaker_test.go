package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShakerSorter_Sort(t *testing.T) {
	tests := []sortTest{
		{
			src:      []Comparable{},
			expected: []Comparable{},
		},
		{
			src:      []Comparable{intC(1)},
			expected: []Comparable{intC(1)},
		},
		{
			src:      []Comparable{intC(1), intC(2)},
			expected: []Comparable{intC(1), intC(2)},
		},
		{
			src:      []Comparable{intC(3), intC(2), intC(1)},
			expected: []Comparable{intC(1), intC(2), intC(3)},
		},
		{
			src:      []Comparable{intC(4), intC(1), intC(10), intC(5), intC(3), intC(8)},
			expected: []Comparable{intC(1), intC(3), intC(4), intC(5), intC(8), intC(10)},
		},
	}

	for i, tt := range tests {
		tt := tt // capture
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			t.Parallel()

			b := shakerSorter{
				src: tt.src,
			}

			b.Sort()

			if !reflect.DeepEqual(tt.src, tt.expected) {
				t.Errorf("Expected to get\n[%#v]\n, but got\n[%#v].", tt.expected, tt.src)
			}
		})
	}
}

func BenchmarkShakerSort_Sort1K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShakerSorter(slice1K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}

func BenchmarkShakerSort_Sort64K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShakerSorter(slice64K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}

func BenchmarkShakerSort_AlmostSorted1K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShakerSorter(almostSorted1K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}

func BenchmarkShakerSort_AlmostSorted64K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShakerSorter(almostSorted64K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}
