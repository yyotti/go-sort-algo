package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/yyotti/go-sort-algo/sort"
)

const randMax = 1000

func main() {
	slice := randomSlice(100)
	fmt.Println(slice)
}

// IntC is `int64` implemented Comparable interface
type IntC int64

// IsLessThan returns i < j
func (i IntC) IsLessThan(j sort.Comparable) bool {
	switch j.(type) {
	case IntC:
		return i < j.(IntC)
	default:
		panic("j is not IntC")
	}

}

func randomSlice(n int) []sort.Comparable {
	var max = big.NewInt(randMax)

	slice := make([]sort.Comparable, 0, n)
	for i := 0; i < n; i++ {
		j, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		slice = append(slice, IntC(j.Int64()))
	}

	return slice
}
