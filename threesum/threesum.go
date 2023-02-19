package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(input)
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	resultMap := make(map[string]struct{})
	totalNums := len(nums)

	// var wg sync.WaitGroup
	// wg.Add(1)
	// defer wg.Done()

	//messages := make(chan string)

	for i := 0; i < totalNums; i++ {
		for j := 1; j < totalNums && i != 1; j++ {
			for k := 2; k < totalNums && j != 2; k++ {
				isDifferentIndices := i != j && i != k && j != k
				isSumZero := nums[i]+nums[j]+nums[k] == 0
				if isDifferentIndices && isSumZero {
					sortedSlice := []int{nums[i], nums[j], nums[k]}
					sort.Ints(sortedSlice)
					strRep := strconv.Itoa(sortedSlice[0]) + "," + strconv.Itoa(sortedSlice[1]) + "," + strconv.Itoa(sortedSlice[2])

					if _, ok := resultMap[strRep]; !ok {
						result = append(result, sortedSlice)
						resultMap[strRep] = struct{}{}
					}
				}
			}
		}
	}

	return result
}
