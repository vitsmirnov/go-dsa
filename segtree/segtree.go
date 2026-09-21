// Segment Tree
package segtree

type InitNode[T any, E any] func(val E) T
type BuildNode[T any] func(leftChild, rightChild T, left, mid, right int) T

type SegTree[T any, E any] struct {
	nodes     []T
	initNode  InitNode[T, E]
	buildNode BuildNode[T]
	size      int
}

func New[T any, E any](
	data []E,
	initNode InitNode[T, E],
	buildNode BuildNode[T]) *SegTree[T, E] {

	size := len(data)
	st := &SegTree[T, E]{
		nodes:     make([]T, size*4),
		initNode:  initNode,
		buildNode: buildNode,
		size:      size}
	st.build(data, 0, 0, size-1)
	return st
}

func (st *SegTree[T, E]) build(data []E, pos int, left, right int) {
	if left == right {
		st.nodes[pos] = st.initNode(data[left])
		return
	}

	mid := left + (right-left)/2
	leftChild, rightChild := pos*2+1, pos*2+2
	st.build(data, leftChild, left, mid)
	st.build(data, rightChild, mid+1, right)
	st.nodes[pos] = st.buildNode(
		st.nodes[leftChild], st.nodes[rightChild],
		left, mid, right)
}

func (st *SegTree[T, E]) Query(left, right int) T {
	return st.query(0, 0, st.size-1, left, right)
}

func (st *SegTree[T, E]) query(pos int, segLeft, segRight, qLeft, qRight int) T {
	if segLeft == qLeft && segRight == qRight {
		return st.nodes[pos]
	}

	mid := segLeft + (segRight-segLeft)/2
	if qRight <= mid {
		return st.query(pos*2+1, segLeft, mid, qLeft, qRight)
	} else if qLeft > mid {
		return st.query(pos*2+2, mid+1, segRight, qLeft, qRight)
	} else {
		resLeft := st.query(pos*2+1, segLeft, mid, qLeft, mid)
		resRight := st.query(pos*2+2, mid+1, segRight, mid+1, qRight)
		return st.buildNode(resLeft, resRight, segLeft, mid, segRight)
	}
}

func (st *SegTree[T, E]) Update(index int, value E) {
	st.update(0, 0, st.size-1, index, value)
}

func (st *SegTree[T, E]) update(pos int, left, right int, index int, value E) {
	if left == right {
		st.nodes[pos] = st.initNode(value)
		return
	}

	mid := left + (right-left)/2
	leftChild, rightChild := pos*2+1, pos*2+2
	if index <= mid {
		st.update(leftChild, left, mid, index, value)
	} else {
		st.update(rightChild, mid+1, right, index, value)
	}
	st.nodes[pos] = st.buildNode(
		st.nodes[leftChild], st.nodes[rightChild],
		left, mid, right)
}

func (st *SegTree[T, E]) Items(key func(node T) E) []E {
	res := make([]E, st.size)
	for i := range res {
		res[i] = key(st.Query(i, i))
	}
	return res
}

func (st *SegTree[T, E]) Size() int {
	return st.size
}
