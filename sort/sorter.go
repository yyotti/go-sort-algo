package sort

// Comparable interface
type Comparable interface {
	IsLessThan(c Comparable) bool
}

// SorterFunc function
type SorterFunc func([]Comparable) []Comparable
