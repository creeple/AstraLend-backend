/**
 * @Author: lixiang
 * @Date: 2025-03-22 14:47
 * @Description: strings.go
 */

package utils

import "strconv"

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}
