package p00376

// WiggleSubsequence ...
func WiggleSubsequence(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	zeroNum := 0
	wiggleNums := []int{}
	for i := 0; i < length-1; i++ {
		wiggleNums = append(wiggleNums, nums[i+1]-nums[i])
		if nums[i+1]-nums[i] == 0 {
			zeroNum++
		}
	}
	if zeroNum == len(wiggleNums) {
		return 1
	}
	// 符号反转次数
	reversalNum := 1
	start, end := 0, 1
	for end < len(wiggleNums) {
		if wiggleNums[start] == 0 {
			end++
			start++
		} else if wiggleNums[end] == 0 {
			end++
		} else if (wiggleNums[start]) > 0 == (wiggleNums[end] > 0) {
			end++
		} else {
			// 符号反转一次，结果就加一
			start = end
			end++
			reversalNum++
		}
	}
	return reversalNum + 1
}
