package exercise1


// Imagine you have been given an array of integers, and a query number k. Your task is to write a function that finds all the triplets in the array that sum up to the query number k.

// Let’s illustrate this with some examples:


// nums = [1, 2, 3, 4, 5]
// count_triplets(nums, k) -> len(triplets)

// to 1
// [1, 2, 3]
// [1, 3, 4]
// [1, 4, 5]
// [1, 5, 2]


// to 2
// [2, 3, 4]
// [2, 4, 5]


// to 3
// [3, 4, 5]
// [3, 5, 1]
// [3, 5, 2]

// to 4
// [4, 5, 1]



func countTriplets(nums []int, k int) int {
	equalK := 0
	for i := 0; i < len(nums); i++ {
		sumTriplets := 0
		for j := 1; j < 3; j++ {
			startPoint := i
			if (i + j) > len(nums) - 1 {
				
			}



			sumTriplets += nums[i] + nums[i + j]
		} 
	}
}
