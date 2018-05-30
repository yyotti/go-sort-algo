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

var slice1K []Comparable
var slice64K []Comparable

var almostSorted1K []Comparable
var almostSorted64K []Comparable

func init() {
	slice1K = make([]Comparable, 1<<10)
	almostSorted1K = make([]Comparable, 1<<10)
	for i := 0; i < len(slice1K); i++ {
		slice1K[i] = intC(i ^ 0x2cc)
		almostSorted1K[i] = intC(i)
	}

	slice64K = make([]Comparable, 1<<16)
	almostSorted64K = make([]Comparable, 1<<16)
	for i := 0; i < len(slice64K); i++ {
		slice64K[i] = intC(i ^ 0xcccc)
		almostSorted64K[i] = intC(i)
	}

	for i := 0; i < len(almostSorted1K)/2; i += len(almostSorted1K) / (1 << 5) {
		almostSorted1K[i], almostSorted1K[len(almostSorted1K)-i-1] = almostSorted1K[len(almostSorted1K)-i-1], almostSorted1K[i]
	}

	for i := 0; i < len(almostSorted64K)/2; i += len(almostSorted64K) / (1 << 8) {
		almostSorted64K[i], almostSorted64K[len(almostSorted64K)-i-1] = almostSorted64K[len(almostSorted64K)-i-1], almostSorted64K[i]
	}
}
