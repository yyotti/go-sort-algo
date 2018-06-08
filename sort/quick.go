package sort

// QuickSort returns slice sorted by `Quick sort`
// https://ja.wikipedia.org/wiki/%E3%82%AF%E3%82%A4%E3%83%83%E3%82%AF%E3%82%BD%E3%83%BC%E3%83%88
func QuickSort(src []Comparable) []Comparable {
	l := len(src)
	if l < 2 {
		return src
	}

	// ピボットには真ん中にある値を選択する
	m := src[(l-1)/2]

	left := make([]Comparable, 0, l)
	right := make([]Comparable, 0, l)
	for _, c := range src {
		if c.IsLessThan(m) {
			left = append(left, c)
		} else if m.IsLessThan(c) {
			right = append(right, c)
		}
	}

	left = append(QuickSort(left), m)
	return append(left, QuickSort(right)...)
}
