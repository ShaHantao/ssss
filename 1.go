package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	a = "Hello world"
	b = "Hello"
	fmt.Println(strings.Index(a, "l"))           //判断a中l第一次出现的位置
	fmt.Println(strings.Index(a, b))             //判断b中任意东西在a中第一次出现的位置
	fmt.Println(strings.LastIndex(a, "l"))       //判断a中l最后一次出现的位置
	fmt.Println(strings.LastIndexAny(a, b))      //判断b中任意东西在a中最后一次出现的位置
	fmt.Println(strings.Contains(a, "e"))        //判断e是否在a中
	fmt.Println(strings.ContainsAny(a, b))       //判断b中的任意东西是否在a中出现
	fmt.Println(strings.EqualFold(a, b))         //不区分大小写，判断a是否与b相等
	fmt.Printf("%q\n", strings.Split(a, " "))    //在空格处将a分成两个子字符串
	fmt.Println(strings.Replace(a, "o", "a", 1)) //从头开始把字符串a中的1个o换成a
	fmt.Println(strings.HasPrefix(a, "hello"))   //判断字符串a是否以hello开头
	fmt.Println(strings.HasPrefix(a, b))         //判断字符串a是否以字符串b开头
}
