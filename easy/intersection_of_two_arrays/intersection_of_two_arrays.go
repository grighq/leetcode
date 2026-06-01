package intersectionoftwoarrays

import "slices"

func intersection(nums1, nums2 []int) []int {
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
			if len(res) == 0 || res[len(res)-1] != nums1[p1] {
				res = append(res, nums1[p1])
			}
			p1++
			p2++
		}
	}

	return res
}

// func intersection(nums1, nums2 []int) []int {
// 	res := []int{}
// 	m := make(map[int]struct{}, len(nums1))
//
// 	for _, num := range nums1 {
// 		m[num] = struct{}{}
// 	}
//
// 	for _, num := range nums2 {
// 		if _, ok := m[num]; ok {
// 			res = append(res, num)
// 			delete(m, num)
// 		}
// 	}
//
// 	return res
// }
