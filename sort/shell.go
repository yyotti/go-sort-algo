package sort

// ShellSort returns slice sorted by `Shell sort`
// https://ja.wikipedia.org/wiki/%E3%82%B7%E3%82%A7%E3%83%AB%E3%82%BD%E3%83%BC%E3%83%88
func ShellSort(src []Comparable) []Comparable {
	l := len(src)

	for h := l / 2; h > 0; h /= 2 {
		for i := h; i < l; i += h {
			e := src[i]
			if e.IsLessThan(src[i-h]) {
				j := i
				for ; j > 0 && e.IsLessThan(src[j-h]); j -= h {
					src[j] = src[j-h]
				}
				src[j] = e
			}
		}
	}

	return src
}
