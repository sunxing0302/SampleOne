package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc axc aaa acb 888 a9c tac"
	//1) 解释规则，它会解析正则表达式，如果成功返回解析器
	//reg1 := regexp.MustCompile(`a.c`) //.表示任意数字
	reg1 := regexp.MustCompile(`a[0-9]c`) //表示首字母是a，总是是0-9的数字，第三个字母是c的表达式
	if reg1 == nil {                      //解析失败，返回nil
		fmt.Println("regexp error")
		return
	}

	//2) 根据规则提取关键信息
	result1 := reg1.FindAllString(buf, -1)
	fmt.Println("result1:", result1)

}
