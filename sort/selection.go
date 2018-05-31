package sort

// SelectionSort returns slice sorted by `Selection sort`
// https://ja.wikipedia.org/wiki/%E9%81%B8%E6%8A%9E%E3%82%BD%E3%83%BC%E3%83%88
func SelectionSort(src []Comparable) []Comparable {
	l := len(src)

	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if src[j].IsLessThan(src[min]) {
				min = j
			}
		}

		src[i], src[min] = src[min], src[i]
	}

	return src
}
