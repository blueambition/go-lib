package str

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

//范围随机数
func RangeRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}

//随机数字码
func RandNumCode(codeLen int) string {
	nums := ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < codeLen; i++ {
		t := rand.Intn(9)
		nums += strconv.Itoa(t)
	}
	return nums
}

//随机数字字符串码
func RandMixCode(codeLen int) string {
	rand.Seed(time.Now().Unix())
	mixArr := [36]string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	mixLen := len(mixArr) - 1
	codes := ""
	for i := 0; i < codeLen; i++ {
		t := rand.Intn(mixLen)
		codes += mixArr[t]
	}
	return codes
}

//任何字符串（中英文）按一个个算长度
func Len(str string) int {
	runes := []rune(str)
	return len(runes)
}

//截断文本
func SubStr(str string, begin int, end int) string {
	if begin < 0 || begin > end {
		return str
	}
	runes := []rune(str)
	if len(runes) >= end {
		runes = runes[begin:end]
		return string(runes)
	}
	return str
}

//截断文本
func Ellipsis(str string, shortLen int) string {
	runes := []rune(str)
	if len(runes) > shortLen {
		runes = runes[:shortLen+1]
		return string(runes) + "......"
	}
	return string(runes)
}

//获取MD5
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//截断文本
func SecretMix(str string, begin, end int, mark string) string {
	mix := ""
	runes := []rune(str)
	strLen := len(runes)
	if strLen < end {
		end = strLen - 1
	}
	if strLen > begin {
		mix = string(runes[0:begin])
		for k, _ := range runes {
			if k >= begin && k <= end {
				mix += mark
			}
		}
		if end < strLen-1 {
			mix += string(runes[end+1:])
		}
	}
	if mix == "" {
		mix = str
	}
	return mix
}
