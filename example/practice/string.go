package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "yym"
	str1 := "中国"

	// len(str) 获取长度
	fmt.Println("str的长度", len(str))   // 3
	fmt.Println("str1的长度", len(str1)) // 1 汉字(utf-8) = 3 字节

	// 字符串遍历, 处理有中文的问题 r := []rune(str)
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符%d=%v\n", i, string(r[i]))
	}

	// 字符串转整数 n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果", n)
	}

	// 整数转字符串
	str2 := strconv.Itoa(12345)
	fmt.Println("整数转字符串", str2)

	// 字符串转字节
	var bytes = []byte("hello")
	fmt.Printf("bytes=%v\n", bytes)

	// byte 转字符串
	str3 := string([]byte{97, 98, 99})
	fmt.Println("str3=\n", str3)

	// 10进制 转 2 8 16 进制
	str4 := strconv.FormatInt(123, 2)
	fmt.Println("str4的二进制", str4)

	// 字符串是否在指定的字符串
	fmt.Println("子串在字符串中吗?=", strings.Contains("yym", "y"))

	// 一个字符串有几个指定的子串
	fmt.Println("有几个子串?=", strings.Count("cheese", "e"))

	// 不区分大小写字符串比较
	fmt.Println("大小写", strings.EqualFold("abc", "Abc"))

	// 返回子串在字符串第一次出现的 index 值, 没有-1
	fmt.Println("出现的下标", strings.Index("NLT_abc", "abc"))

	// 返回子串在字符串最后一次出现的 index 值, 没有-1
	fmt.Println("最后一次出现的下标", strings.LastIndex("go golang", "go"))

	// 子串替换
	fmt.Println("子串替换", strings.Replace("go go golang", "go", "go语言", 2))

	// 使用指定分隔符, 分割字符串
	fmt.Println("分割字符串", strings.Split("gogogolang", "o")) //  [g g g lang]

	// 将字符串的字母进行大小写替换
	fmt.Println("字符串大小写替换", strings.ToLower("Go"), strings.ToUpper("Go"))

	// 去掉字符传两边的空格
	fmt.Println("字符串两边的空格", strings.TrimSpace("! hello ! "))

	// 字符串两边指定的字符去掉
	fmt.Println("两边指定字符", strings.Trim("! hello!", "!"))

	fmt.Println("左边指定字符", strings.TrimLeft("! hello!", "!"))
}
