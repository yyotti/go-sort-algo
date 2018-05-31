package sort

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	test(t, ShellSort)
}

func BenchmarkShellSort_Sort1K(b *testing.B) {
	bench(b, ShellSort, slice1K)
}

func BenchmarkShellSort_Sort16K(b *testing.B) {
	bench(b, ShellSort, slice16K)
}

func BenchmarkShellSort_AlmostSorted1K(b *testing.B) {
	bench(b, ShellSort, almostSorted1K)
}

func BenchmarkShellSort_AlmostSorted16K(b *testing.B) {
	bench(b, ShellSort, almostSorted16K)
}
