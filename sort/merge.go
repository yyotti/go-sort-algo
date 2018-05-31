package sort

// MergeSort returns slice sorted by `Merge sort`
// https://ja.wikipedia.org/wiki/%E3%83%9E%E3%83%BC%E3%82%B8%E3%82%BD%E3%83%BC%E3%83%88
func MergeSort(src []Comparable) []Comparable {
	l := len(src)
	if l < 2 {
		return src
	}

	half1 := MergeSort(src[:l/2])
	half2 := MergeSort(src[l/2:])

	dest := make([]Comparable, 0, l)
	i, j := 0, 0
	for i < len(half1) && j < len(half2) {
		if half1[i].IsLessThan(half2[j]) {
			dest = append(dest, half1[i])
			i++
		} else {
			dest = append(dest, half2[j])
			j++
		}
	}

	for ; i < len(half1); i++ {
		dest = append(dest, half1[i])
	}
	for ; j < len(half2); j++ {
		dest = append(dest, half2[j])
	}

	return dest
}
