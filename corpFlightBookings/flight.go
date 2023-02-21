package main

import "fmt"

func main() {
	//[10,55,45,25,25]
	input := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	result := corpFlightBookings(input, 5)
	fmt.Println(result)
}

// func corpFlightBookings(bookings [][]int, n int) []int {
// 	result := make([]int, n)
// 	for i := 0; i < len(bookings); i++ {
// 		for start := bookings[i][0]; start <= bookings[i][1]; start++ {
// 			result[start-1] += bookings[i][2]
// 		}
// 	}
// 	return result
// }

//Brute force
// func corpFlightBookings(bookings [][]int, n int) []int {
// 	result := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < len(bookings); j++ {
// 			if bookings[j][0] <= i+1 && i+1 <= bookings[j][1] {
// 				result[i] = result[i] + bookings[j][2]
// 			}
// 		}
// 	}
// 	return result
// }

func corpFlightBookings(bookings [][]int, n int) []int {

	// +1 as dummy guard on the tail, which allow us not to check right boundary every time
	unitStep := make([]int, n+1)

	for _, booking := range bookings {

		// -1 because booking flight is 1-indexed, given by description
		left, right, seatVector := booking[0]-1, booking[1]-1, booking[2]

		unitStep[left] += seatVector
		unitStep[right+1] -= seatVector
	}

	// Reconstruct booking as drawing combination of retangle signal, built with unit step impulse
	for i := 1; i < len(unitStep); i++ {
		unitStep[i] += unitStep[i-1]
	}

	// last one is dummy guard on the tail, no need to return
	return unitStep[0:n]
}
