package p00033

// SearchInRotatedSortedArray ...
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	firstValue := nums[0]
	table := [8]bool{false, true, false, false, true, true, false, true}
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		var tp, mp, mt int
		if target < firstValue {
			// 目标值在右半部分
			tp = 1 << 2
		}
		if nums[mid] < firstValue {
			// 中间值在右半部分
			mp = 1 << 1
		}
		if target > nums[mid] {
			// 目标值大于中间值
			mt = 1
		}
		if table[tp+mp+mt] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
