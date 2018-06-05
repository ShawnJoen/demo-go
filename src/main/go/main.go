package main

import (
	"fmt"
)

/*
声明变量:
1. var 名称 类型 	#未赋值为默认值
   名称 = 赋值
2. var 名称 = 赋值 	#自动类型推断
3. var (名称 类型) 	#一般用于全局变量
4. var d, e int = 2, 4
*/
var a string = "shawn.com"
var b bool
var c = "学习笔记"
var (
	gv1 int
	gv2 int8
)
//多变量声明
var d, e int = 2, 4
var f, g = "3", "shawn"

func main() {
	fmt.Println("Hello, World!")//Hello, World!
	fmt.Println(a, b, c)//shawn.com false 学习笔记
	fmt.Println("多变量声明:", d, e, f, g)//多变量声明: 2 4 3 shawn
}
