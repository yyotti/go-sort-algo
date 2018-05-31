package sort

// BubbleSort returns slice sorted by `Bubble sort`
// https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%96%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88
func BubbleSort(src []Comparable) []Comparable {
	l := len(src)

	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if src[j].IsLessThan(src[j-1]) {
				src[j], src[j-1] = src[j-1], src[j]
			}
		}
	}

	return src
}
