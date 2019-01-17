package p00074

// SearchA2DMatrix ...
func SearchA2DMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	// 每一行的长度
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	low, high := 0, row*col-1
	for low <= high {
		mid := (low + high) / 2
		midValue := matrix[mid/col][mid%col]
		if target > midValue {
			low = mid + 1
		} else if target < midValue {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
