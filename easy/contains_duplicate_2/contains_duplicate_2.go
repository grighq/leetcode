package containsduplicate2

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[num]; ok && i-j <= k {
			return true
		}
		m[num] = i
	}

	return false
}

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	window := make(map[int]struct{})
// 	for i, num := range nums {
// 		if _, ok := window[num]; ok {
// 			return true
// 		}
//
// 		window[num] = struct{}{}
//
// 		if len(window) > k {
// 			delete(window, nums[i-k])
// 		}
// 	}
//
// 	return false
// }
