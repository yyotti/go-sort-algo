package sort

import (
	"fmt"
	"reflect"
	"testing"
)

type sortTest struct {
	src      []Comparable
	expected []Comparable
}

type intC int64

func (i intC) IsLessThan(j Comparable) bool {
	switch j.(type) {
	case intC:
		return i < j.(intC)
	default:
		panic("j is not intC")
	}

}

var slice1K []Comparable
var slice16K []Comparable

var almostSorted1K []Comparable
var almostSorted16K []Comparable

func init() {
	slice1K = make([]Comparable, 1<<10)
	almostSorted1K = make([]Comparable, 1<<10)
	for i := 0; i < len(slice1K); i++ {
		slice1K[i] = intC(i ^ 0x2cc)
		almostSorted1K[i] = intC(i)
	}

	slice16K = make([]Comparable, 1<<14)
	almostSorted16K = make([]Comparable, 1<<14)
	for i := 0; i < len(slice16K); i++ {
		slice16K[i] = intC(i ^ 0xcccc)
		almostSorted16K[i] = intC(i)
	}

	for i := 0; i < len(almostSorted1K)/2; i += len(almostSorted1K) / (1 << 5) {
		almostSorted1K[i], almostSorted1K[len(almostSorted1K)-i-1] = almostSorted1K[len(almostSorted1K)-i-1], almostSorted1K[i]
	}

	for i := 0; i < len(almostSorted16K)/2; i += len(almostSorted16K) / (1 << 8) {
		almostSorted16K[i], almostSorted16K[len(almostSorted16K)-i-1] = almostSorted16K[len(almostSorted16K)-i-1], almostSorted16K[i]
	}
}

func test(t *testing.T, algo SorterFunc) {
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
			src:      []Comparable{intC(2), intC(1)},
			expected: []Comparable{intC(1), intC(2)},
		},
		{
			src:      []Comparable{intC(1), intC(2), intC(3)},
			expected: []Comparable{intC(1), intC(2), intC(3)},
		},
		{
			src:      []Comparable{intC(3), intC(1), intC(2)},
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

			actual := algo(tt.src)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected to get\n[%#v]\n, but got\n[%#v].", tt.expected, tt.src)
			}
		})
	}
}

func bench(b *testing.B, algo SorterFunc, slice []Comparable) {
	newSlice := make([]Comparable, len(slice))

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		copy(newSlice, slice)

		b.StartTimer()
		algo(newSlice)
		b.StopTimer()
	}
}
