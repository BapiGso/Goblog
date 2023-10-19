package mymiddleware

import "strconv"

// ValidateNum 首页返回1，不是数字返回err调用404，其他为对应页数
func ValidateNum(numStr string) (int, error) {
	if numStr == "" {
		return 1, nil
	}
	num, err := strconv.ParseUint(numStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}
