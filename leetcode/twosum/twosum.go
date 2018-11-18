/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

package twosum

// TwoSum -- Use map to save intermediate results, the key is 'target - num', the value is index of num in nums.
func TwoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			// not found the second one
			m[target-num] = i
		} else {
			// found the second one
			result = append(result, m[num])
			result = append(result, i)
			break
		}
	}
	return result
}
