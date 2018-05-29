package sort

// Comparable interface
type Comparable interface {
	IsLessThan(c Comparable) bool
}

// Sorter interface
type Sorter interface {
	Sort()
}
