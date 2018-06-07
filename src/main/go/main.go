package main

import (
	"fmt"
	"unsafe"//常量函数
)

/*
声明变量:
1. var 名称 类型 	#未赋值为默认值
   名称 = 赋值
2. var 名称 = 赋值 	#自动类型推断
3. var (名称 类型) 	#一般用于全局变量
4. var d, e int = 2, 4
5. 只能被用在函数体内而不可以用于全局变量的声明与赋值
	相同的代码块中不可以再次对于相同名称的变量使用初始化声明a := 20.但是 a = 20是可以因为这是给相同的变量赋予一个新的值

变量使用: 全局变量是允许声明后不使用 不过函数体内的变量相同的代码块中必须使用它不然会编译错误
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

/*func main() {
	fmt.Println("Hello, World!")//Hello, World!
	h := 77//只能被用在函数体内,不可以用于全局变量的声明与赋值且必须使用
	fmt.Println(a, b, c, h)//shawn.com false 学习笔记 77
	fmt.Println("多变量声明:", d, e, f, g)//多变量声明: 2 4 3 shawn
}*/

/*
声明常量:
1. const b string = "abc" 显式类型定义
2. const b = "abc" 隐式类型定义
3. const c_name1, c_name2 = value1, value2 多个相同类型的声明
4. const (
		CONST4 = 0
		CONST5 = 1
		CONST6 = 2
	)枚举
5. iota，特殊常量，可以认为是一个可以被编译器修改的常量
	每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
常量函数 len(), cap(), unsafe.Sizeof()函数计算表达式的值
unsafe.Sizeof:
	字符串类型在go里是个结构,包含指向底层数组的指针和长度,这两部分每部分都是8个字节，
	所以字符串类型大小为16个字节
*/

const (
	G_CONST1 = "abc"
	G_CONST2 = len(G_CONST1)
	G_CONST3 = unsafe.Sizeof(G_CONST1)//获取占用空间大小
)
const (
	G_CONST_IOTA1 = iota
	G_CONST_IOTA2 = iota
	G_CONST_IOTA3 = iota
)
func main() {
	const LENGTH int = 10
	const WIDTH = 8
	const CONST1, CONST2, CONST3 = 1, false, "str"
	const (
		CONST4 = 0
		CONST5 = 1
		CONST6 = 2
	)
	fmt.Printf("面积为 : %d", LENGTH * WIDTH)//面积为 : 80
	println()
	println(CONST1, CONST2, CONST3, CONST4, CONST5, CONST6)//1 false str 0 1 2
	println(G_CONST1, G_CONST2, G_CONST3)//abc 3 16
	println(G_CONST_IOTA1, G_CONST_IOTA2, G_CONST_IOTA3)//0 1 2
	const (
		IOTA1 = iota   //0
		IOTA2          //1
		IOTA3          //2
		IOTA4 = "shawn"//shawn, iota += 1
		IOTA5          //shawn, iota += 1
		IOTA6 = 100    //iota += 1
		IOTA7 = iota   //6, 恢复计数
		IOTA8          //7
	)
	println(IOTA1, IOTA2, IOTA3, IOTA4, IOTA5, IOTA6, IOTA7, IOTA8);//0 1 2 shawn shawn 100 6 7
	const (
		IOTA9 = iota   //0
		IOTA10         //1
	)
	println(IOTA9, IOTA10)//0 1
/*
	位运算符:
	假定 A = 60; B = 13; 其二进制数转换为：
	A = 0011 1100
	B = 0000 1101
	-----------------
	A&B = 0000 1100
	A|B = 0011 1101
	A^B = 0011 0001
*/
	var va uint = 60    /* 60 = 0011 1100 */
	var vb uint = 13    /* 13 = 0000 1101 */
	var vc uint = 0
	vc = va & vb       /* 12 = 0000 1100 */
	fmt.Printf("第一行 - vc 的值为 %d\n", vc )//第一行 - vc 的值为 12
	vc = va | vb       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - vc 的值为 %d\n", vc )//第二行 - vc 的值为 61
	vc = va ^ vb       /* 49 = 0011 0001 */
	fmt.Printf("第三行 - vc 的值为 %d\n", vc )//第三行 - vc 的值为 49
	/*
	左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，
	由"<<"右边的数指定移动的位数，高位丢弃，低位补0
	*/
	vc = va << 2     /* 240 = 1111 0000 */
	fmt.Printf("第四行 - vc 的值为 %d\n", vc )//第四行 - vc 的值为 240
	/*
	右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位
	全部右移若干位，">>"右边的数指定移动的位数
	*/
	vc = va >> 2     /* 15 = 0000 1111 */
	fmt.Printf("第五行 - vc 的值为 %d\n", vc )//第五行 - vc 的值为 15
/*
其他运算符: 指针
&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量			*a; 是一个指针变量
*/
	var a int = 4
	var b int32
	var c float32
	var ptr *int
	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );//第 1 行 - a 变量类型为 = int
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );//第 2 行 - b 变量类型为 = int32
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );//第 3 行 - c 变量类型为 = float32
	/*  & 和 * 运算符实例 */
	ptr = &a    /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a);//a 的值为  4
	fmt.Printf("a 的地址为  %d\n", ptr);//a 的地址为  825741361352
	fmt.Printf("*ptr 为 %d\n", *ptr);//*ptr 为 4
}


































