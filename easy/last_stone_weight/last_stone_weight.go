package laststoneweight

import "slices"

func lastStoneWeight(stones []int) int {
	slices.Sort(stones)
	return helper(stones)
}

func helper(stones []int) int {
	count := len(stones)
	switch count {
	case 0:
		return 0
	case 1:
		return stones[0]
	}

	smash := stones[count-1] - stones[count-2]
	stones = stones[:count-2]

	if smash > 0 {
		idx, _ := slices.BinarySearch(stones, smash)
		stones = append(stones, 0)
		copy(stones[idx+1:], stones[idx:])
		stones[idx] = smash
	}

	return helper(stones)
}

// type IntHeap []int
//
// func (h IntHeap) Len() int { return len(h) }
//
// func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
// func (h *IntHeap) Push(x any) {
// 	*h = append(*h, x.(int))
// }
//
// func (h *IntHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }
//
// func lastStoneWeight(stones []int) int {
// 	h := (*IntHeap)(&stones)
// 	heap.Init(h)
//
// 	for h.Len() > 1 {
// 		first := heap.Pop(h).(int)
// 		second := heap.Pop(h).(int)
//
// 		if first != second {
// 			heap.Push(h, first-second)
// 		}
// 	}
//
// 	if h.Len() == 0 {
// 		return 0
// 	}
// 	return heap.Pop(h).(int)
// }
