package sort

// CombSort returns slice sorted by `Comb sort`
// https://ja.wikipedia.org/wiki/%E3%82%B3%E3%83%A0%E3%82%BD%E3%83%BC%E3%83%88
func CombSort(src []Comparable) []Comparable {
	l := len(src)
	h := l

	var swapped bool
	for h > 1 || swapped {
		if h > 1 {
			h = h * 10 / 13
		}

		swapped = false
		for i := 0; i+h < l; i++ {
			if src[i+h].IsLessThan(src[i]) {
				src[i], src[i+h] = src[i+h], src[i]
				swapped = true
			}
		}
	}

	return src
}
