package sort

// ParallelMergeSort returns slice sorted by `ParallelMerge sort`
// https://ja.wikipedia.org/wiki/%E3%83%9E%E3%83%BC%E3%82%B8%E3%82%BD%E3%83%BC%E3%83%88
func ParallelMergeSort(src []Comparable) []Comparable {
	l := len(src)
	if l < 2 {
		return src
	}

	ch := make(chan []Comparable, 2)
	go func(src []Comparable, ch chan<- []Comparable) {
		ch <- ParallelMergeSort(src)
	}(src[:l/2], ch)
	go func(src []Comparable, ch chan<- []Comparable) {
		ch <- ParallelMergeSort(src)
	}(src[l/2:], ch)

	half1 := <-ch
	half2 := <-ch

	close(ch)

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
