package p0001

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
