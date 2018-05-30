package sort

type bubbleSorter struct {
	src []Comparable
}

// Sort returns slice sorted by `Bubble sort`
// https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%96%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88
func (b bubbleSorter) Sort() {
	src := b.src
	l := len(src)

	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if src[j].IsLessThan(src[j-1]) {
				src[j], src[j-1] = src[j-1], src[j]
			}
		}
	}
}

// NewBubbleSorter returns bubbleSorter instance
func NewBubbleSorter(src []Comparable) Sorter {
	s := make([]Comparable, 0, len(src))
	copy(s, src)

	return bubbleSorter{
		src: s,
	}
}
