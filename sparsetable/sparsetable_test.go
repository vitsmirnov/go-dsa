package sparsetable

import (
	"math/rand/v2"
	"testing"
)

func TestSparseTableIdempQuery(t *testing.T) {
	const minNum int = -1e5
	const maxNum int = 1e5
	const numsRange = maxNum - minNum + 1
	const numsLen int = 1e5
	const testCount int = 100
	const queryCount int = 100
	funcs := []func(int, int) int{
		func(a, b int) int { return min(a, b) },
		func(a, b int) int { return max(a, b) },
		func(a, b int) int { return a | b }}
	nums := make([]int, numsLen)
	for range testCount {
		for i := range nums {
			nums[i] = rand.IntN(numsRange) + minNum
		}
		f := funcs[rand.IntN(len(funcs))]
		st := New(nums, f)
		if st.Size() != len(nums) {
			t.Errorf("sparse table size incorrect: %v (%v expected)", st.Size(), len(nums))
		}
		for range queryCount {
			length := rand.IntN(numsLen-1) + 1
			left := rand.IntN(numsLen)
			right := min(left+length-1, numsLen-1)
			res1 := st.IdempQuery(left, right)
			res2 := agr(nums[left:right+1], f)
			if res1 != res2 {
				t.Errorf("%v != %v: [%v, %v]", res1, res2, left, right)
			}
		}
	}
}

func TestSparseTableNonIdempQuery(t *testing.T) {
	const minNum int = -1e2
	const maxNum int = 1e2
	const numsRange = maxNum - minNum + 1
	const numsLen int = 1e5
	const testCount int = 100
	const queryCount int = 500
	nums := make([]int, numsLen)
	sum := func(a, b int) int { return a + b }
	for range testCount {
		for i := range nums {
			nums[i] = rand.IntN(numsRange) + minNum
		}
		st := New(nums, sum)
		if st.Size() != len(nums) {
			t.Errorf("sparse table size incorrect: %v (%v expected)", st.Size(), len(nums))
		}
		for range queryCount {
			length := rand.IntN(numsLen-1) + 1
			left := rand.IntN(numsLen)
			right := min(left+length-1, numsLen-1)
			res1 := st.NonIdempQuery(left, right)
			res2 := agr(nums[left:right+1], sum)
			if res1 != res2 {
				t.Errorf("%v != %v: [%v, %v]", res1, res2, left, right)
			}
		}
	}
}

func agr(nums []int, f func(int, int) int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		res = f(res, num)
	}
	return res
}
