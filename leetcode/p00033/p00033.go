package p00033

// SearchInRotatedSortedArray ...
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	firstValue := nums[0]
	// 目标值是否在左半部分
	inLeft := target >= firstValue

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if inLeft {
			if nums[mid] >= firstValue {
				if nums[mid] > target {
					end = mid - 1
				} else if nums[mid] < target {
					start = mid + 1
				}
			} else {
				end = mid - 1
			}
		} else {
			if nums[mid] >= firstValue {
				start = mid + 1
			} else {
				if nums[mid] > target {
					end = mid - 1
				} else if nums[mid] < target {
					start = mid + 1
				}
			}
		}
	}
	return -1
}
