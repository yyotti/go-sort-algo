package sort

type shakerSorter struct {
	src []Comparable
}

// Sort returns slice sorted by `Shaker sort`
// https://ja.wikipedia.org/wiki/%E3%82%B7%E3%82%A7%E3%83%BC%E3%82%AB%E3%83%BC%E3%82%BD%E3%83%BC%E3%83%88
func (b shakerSorter) Sort() {
	src := b.src

	top := 0
	bottom := len(src) - 1

	for {
		// 順方向スキャン
		lastSwap := top

		for i := top; i < bottom; i++ {
			if src[i+1].IsLessThan(src[i]) {
				src[i], src[i+1] = src[i+1], src[i]
				lastSwap = i
			}
		}

		// 後方のスキャン範囲を狭める
		bottom = lastSwap

		if top == bottom {
			break
		}

		// 逆方向スキャン
		lastSwap = bottom

		for i := bottom; i > top; i-- {
			if src[i].IsLessThan(src[i-1]) {
				src[i], src[i-1] = src[i-1], src[i]
				lastSwap = i
			}
		}

		// 前方のスキャン範囲を狭める
		top = lastSwap

		if top == bottom {
			break
		}
	}
}

// NewShakerSorter returns shakerSorter instance
func NewShakerSorter(src []Comparable) Sorter {
	s := make([]Comparable, 0, len(src))
	copy(s, src)

	return shakerSorter{
		src: s,
	}
}
