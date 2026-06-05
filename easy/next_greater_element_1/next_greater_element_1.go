package nextgreaterelement1

func nextGreaterElement(nums1, nums2 []int) []int {
	stack := []int{}
	greaters := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		curr := nums2[i]

		for len(stack) > 0 && stack[len(stack)-1] < curr {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			greaters[curr] = stack[len(stack)-1]
		} else {
			greaters[curr] = -1
		}

		stack = append(stack, curr)
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = greaters[num]
	}

	return res
}
