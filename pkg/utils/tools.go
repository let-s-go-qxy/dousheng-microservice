package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// GetMd5Str 根据传入字符串获取MD5加密后的长度32位字符串
func GetMd5Str(str string) string {
	data := []byte(str)
	md5Ret := md5.Sum(data)
	return fmt.Sprintf("%x", md5Ret)
}

// String2Int string数组转int数组
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
