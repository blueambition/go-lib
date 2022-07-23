package str

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
func ShortTxt(str string, shortLen int) string {
	runes := []rune(str)
	if len(runes) > shortLen {
		runes = runes[:shortLen+1]
		return string(runes) + "......"
	}
	return string(runes)
}

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//去除HTML标签
func TrimHtml(src string, breakLine bool) string {
	src = strings.ReplaceAll(src, "&#8211;", "–")
	src = strings.ReplaceAll(src, "&nbsp;", " ")
	src = strings.ReplaceAll(src, "&#160;", " ")
	src = strings.ReplaceAll(src, "&#32;", " ")
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	if breakLine {
		//替换<p>标签
		re, _ = regexp.Compile("<p[^<>]+?>|<p />")
		src = re.ReplaceAllString(src, "\n")
		//替换<br>标签
		re, _ = regexp.Compile("<br[^<>]+?>|<br />")
		src = re.ReplaceAllString(src, "\n")
	}
	//去除STYLE
	re, _ = regexp.Compile("<style[^<>]+?>(.*?)</style>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("<script[^<>]+?></script>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("<[^<>]+?>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	//re, _ = regexp.Compile("\\s+")
	//src = re.ReplaceAllString(src, " ")
	re, _ = regexp.Compile("\n+")
	src = re.ReplaceAllString(src, "\n")
	//特殊符号
	re, _ = regexp.Compile("&.+?;")
	src = re.ReplaceAllString(src, "")
	return src
}

//<br> <p>标签转换成换行
func TagToLine(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//替换<p>标签
	re, _ = regexp.Compile("<p[^<>]+?>|<p />")
	src = re.ReplaceAllString(src, "\n")
	//替换<br>标签
	re, _ = regexp.Compile("<br[^<>]+?>|<br />")
	src = re.ReplaceAllString(src, "\n")
	return src
}

//unicode索引位置
func IndexOf(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

//unicode索引位置
func LastIndexOf(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.LastIndex(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

//去除换行
func TrimLine(html string) string {
	html = strings.Replace(html, "\n", "", -1)
	html = strings.Replace(html, "\t", "", -1)
	return html
}
