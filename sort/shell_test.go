package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShellSorter_Sort(t *testing.T) {
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

			b := shellSorter{
				src: tt.src,
			}

			b.Sort()

			if !reflect.DeepEqual(tt.src, tt.expected) {
				t.Errorf("Expected to get\n[%#v]\n, but got\n[%#v].", tt.expected, tt.src)
			}
		})
	}
}

func BenchmarkShellSort_Sort1K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShellSorter(slice1K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}

func BenchmarkShellSort_Sort64K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShellSorter(slice64K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}

func BenchmarkShellSort_AlmostSorted1K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShellSorter(almostSorted1K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}

func BenchmarkShellSort_AlmostSorted64K(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		sorter := NewShellSorter(almostSorted64K)

		b.StartTimer()
		sorter.Sort()
		b.StopTimer()
	}
}
