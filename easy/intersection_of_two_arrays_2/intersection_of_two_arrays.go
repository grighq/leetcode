package intersectionoftwoarrays2

import "slices"

func intersect(nums1, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	res := []int{}
	p1, p2 := 0, 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] > nums2[p2] {
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			res = append(res, nums1[p1])
			p1++
			p2++
		}
	}

	return res
}

// func intersect(nums1, nums2 []int) []int {
// 	nums := make(map[int]int)
// 	res := []int{}
// 	for _, num := range nums1 {
// 		nums[num]++
// 	}
//
// 	for _, num := range nums2 {
// 		if nums[num] > 0 {
// 			res = append(res, num)
// 			nums[num]--
// 		}
// 	}
//
// 	return res
// }
