package three

import (
	"fmt"
	"sort"
)

func main() {
	input := threeSum([]int{-70, -10, -10, 1, 2, -1, 20})
	fmt.Println(input)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	totalNums := len(nums)

	for i := 0; i < totalNums-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate elements
		}
		target := -nums[i]
		j, k := i+1, totalNums-1
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++ // skip duplicate elements
				}
				for j < k && nums[k] == nums[k+1] {
					k-- // skip duplicate elements
				}
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}

	return result
}
