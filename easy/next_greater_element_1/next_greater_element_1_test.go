package nextgreaterelement1

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2, 0}, []int{3, 4, 2, 0, 1}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}
