package exec

/**
 * Description: 求长度为5的数组 总和 与 平均值
 */
func Average(arr [5]int) (sum int, aver float64) {
	for _, val := range arr {
		sum += val
	}
	aver = float64(sum) / float64(len(arr))
	return
}
