// Sparse Table
package sparsetable

import (
	"math/bits"
)

func NewMinST(nums []int) *SparseTable {
	return New(nums, func(a, b int) int { return min(a, b) })
}
func NewMaxST(nums []int) *SparseTable {
	return New(nums, func(a, b int) int { return max(a, b) })
}
func NewOrST(nums []int) *SparseTable {
	return New(nums, func(a, b int) int { return a | b })
}

type SparseTable struct {
	items [][]int
	f     func(int, int) int
}

func New(nums []int, f func(int, int) int) *SparseTable {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	levelCount := bits.Len(uint(numsLen))
	items := make([][]int, levelCount)
	items[0] = make([]int, numsLen)
	copy(items[0], nums)
	for level := 1; level < levelCount; level++ {
		items[level] = make([]int, numsLen-(1<<level)+1)
		prevLevel := level - 1
		prevLevelLen := 1 << prevLevel
		for i := range items[level] {
			items[level][i] = f(items[prevLevel][i], items[prevLevel][i+prevLevelLen])
		}
	}
	return &SparseTable{
		items: items,
		f:     f}
}

// idempotent query (min/max, or, etc)
func (st *SparseTable) IdempQuery(left, right int) int {
	l := bits.Len(uint(right-left+1)) - 1
	return st.f(st.items[l][left], st.items[l][right-(1<<l)+1])
}

// non-idempotent query (sum, xor etc)
func (st *SparseTable) NonIdempQuery(left, right int) int {
	sum := 0
	length := right - left + 1
	for level := 0; length != 0; level++ {
		if length&1 == 1 {
			sum = st.f(sum, st.items[level][left])
			left += 1 << level
		}
		length >>= 1
	}
	return sum
}
