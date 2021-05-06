package exec

/**
 * Description: 求长度为5的数组的最大值
 */
func Exec02(arr [5]int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
