package sort

// InsertionSort returns slice sorted by `Insertion sort`
// https://ja.wikipedia.org/wiki/%E6%8C%BF%E5%85%A5%E3%82%BD%E3%83%BC%E3%83%88
func InsertionSort(src []Comparable) []Comparable {
	l := len(src)

	for i := 1; i < l; i++ {
		e := src[i]
		if e.IsLessThan(src[i-1]) {
			j := i
			for ; j > 0 && e.IsLessThan(src[j-1]); j-- {
				src[j] = src[j-1]
			}
			src[j] = e
		}
	}

	return src
}
