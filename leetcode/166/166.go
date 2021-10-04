package _66

import (
	"strconv"
)

//166. 分数到小数
func fractionToDecimal(numerator int, denominator int) string {
	if numerator % denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	s := []byte{}
	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}

	numerator = abs(numerator)
	denominator = abs(denominator)
	// 整数部分
	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')

	// 小数部分
	idxMap := map[int]int{}
	remainder := numerator % denominator
	for remainder != 0 && idxMap[remainder] == 0 {
		idxMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0' + byte(remainder / denominator))
		remainder %= denominator
	}
	// 判断是否有循环
	if remainder > 0 {
		insertIdx := idxMap[remainder]
		s = append(s[:insertIdx], append([]byte{'('}, s[insertIdx:]...)...)
		s = append(s, ')')
	}
	return string(s)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}