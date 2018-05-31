package sort

// GnomeSort returns slice sorted by `Gnome sort`
// https://ja.wikipedia.org/wiki/%E3%83%8E%E3%83%BC%E3%83%A0%E3%82%BD%E3%83%BC%E3%83%88
func GnomeSort(src []Comparable) []Comparable {
	l := len(src)

	for i := 1; i < l; {
		if src[i].IsLessThan(src[i-1]) {
			src[i], src[i-1] = src[i-1], src[i]
			i--
			if i == 0 {
				i++
			}
		} else {
			i++
		}
	}

	return src
}
