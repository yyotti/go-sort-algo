package sort

type sortTest struct {
	src      []Comparable
	expected []Comparable
}

type intC int64

func (i intC) IsLessThan(j Comparable) bool {
	switch j.(type) {
	case intC:
		return i < j.(intC)
	default:
		panic("j is not intC")
	}

}
