package p00053

// MaximumValue ...
func MaximumValue(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	result := nums[0]
	f := 0
	for _, num := range nums {
		// 每次只使用 f[k-1]
		temp := f + num
		if num > temp {
			f = num
		} else {
			f = temp
		}
		if f > result {
			result = f
		}
	}
	return result
}

// MaximumSubarray 返回长度为3的数组，下标0为 最大和，下标1，2分别表示子串在nums中的下标起止（闭区间）。
func MaximumSubarray(nums []int) [3]int {
	if len(nums) <= 0 {
		return [3]int{}
	}
	result := nums[0]
	f := 0
	start := 0
	end := 0
	for k, num := range nums {
		// 每次只使用 f[k-1]
		temp := f + num
		if num > temp {
			f = num
			// nums[k - i] == f[k - i] 时开始
			start = k
		} else {
			f = temp
		}
		if f > result {
			result = f
			// k处结束
			end = k
		}
	}
	return [3]int{result, start, end}
}
