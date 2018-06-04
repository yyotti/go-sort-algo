package sort

// HeapSort returns slice sorted by `Heap sort`
// https://ja.wikipedia.org/wiki/%E3%83%92%E3%83%BC%E3%83%97%E3%82%BD%E3%83%BC%E3%83%88
func HeapSort(src []Comparable) []Comparable {
	i := 0
	for i++; i < len(src); i++ {
		upHeap(src, i)
	}

	for i--; i > 0; i-- {
		src[0], src[i] = src[i], src[0]
		downHeap(src, i)
	}

	return src
}

func upHeap(src []Comparable, n int) {
	i := n
	for i > 0 {
		j := (i+1)/2 - 1
		if !src[j].IsLessThan(src[i]) {
			break
		}
		src[j], src[i] = src[i], src[j]
		i = j
	}
}

func downHeap(src []Comparable, n int) {
	i, j := 0, 0

	for {
		l := (i+1)*2 - 1
		r := (i + 1) * 2
		if l >= n {
			break
		}

		if src[j].IsLessThan(src[l]) {
			j = l
		}

		if r < n && src[j].IsLessThan(src[r]) {
			j = r
		}

		if j == i {
			break
		}

		src[j], src[i] = src[i], src[j]
		i = j
	}
}
