package exec

import (
	"math/rand"
	"time"
)

/**
 * Description: 生成一个随机的int数组 大小为5
 */
func RandArr() [5]int {
	var arr [5]int
	// 为了每次生成随机数不同，需要给个seed值
	rand.Seed(time.Now().UnixNano()) // 给其当前的时间戳 给纳秒比较好
	for i, _ := range arr {
		arr[i] = rand.Intn(100) // 生成0 - 100 的随机int数
	}
	return arr
}

/**
 * Description: 翻转数组
随机生成五个数 传入数组 并将其翻转
*/
func Reverse(arr *[5]int) {
	low := 0
	high := len(arr) - 1
	for low < high {
		swap(&(*arr)[low], &(*arr)[high])
		low++
		high--
	}
}

/**
 * Description: 交换两元素的值
 */
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
