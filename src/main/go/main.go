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

/********************************************************************数据类型 浮点数*/
	//浮点数计算输出有一定的偏差，你也可以转整型来设置精度
	v1 := 1690           				//表示1.69
	b1 := 1700           				//表示1.70
	c1 := v1 * b1          				//结果应该是2873000表示 2.873
	fmt.Println("内部编码", c1)  	//内部编码 2873000
	fmt.Println(float64(c1) / 1000000)	//2.873

/********************************************************************条件语句*/
	var var1, var2 = 100, 200
	if var1 == 100 {
		if var2 == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("var1 的值为 100 ， var2 的值为 200\n" )//var1 的值为 100 ， var2 的值为 200
		}
	}
/* switch语句
1. 不需要break
2. 变量必须是相同的类型
3. 多个条件使用逗号分割
*/
	var grade string = "C"
	switch {
		case grade == "A" :
			fmt.Printf("优秀!\n" )
		case grade == "B", grade == "C" :
			fmt.Printf("良好\n" )
		case grade == "D" :
			fmt.Printf("及格\n" )
		case grade == "F":
			fmt.Printf("不及格\n" )
		default:
			fmt.Printf("差\n" )
	}//良好

	var marks int = 70
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"
	}
	fmt.Printf("你的等级是 %s\n", grade )//你的等级是 C

	type TestStruct1 struct {
		Key string
	}
	var x interface{}			//x 的类型 :<nil>
	x = TestStruct1 {Key: "3"} 	//x 是 TestStruct1 struct 型
	x = true 					//x 是 bool 或 string 型
	x = 10						//x 是 int 型
	switch i := x.(type) {
		case nil:
			fmt.Printf(" x 的类型 :%T",i)
		case int:
			fmt.Printf("x 是 int 型")
		case TestStruct1:
			fmt.Printf("x 是 TestStruct1 struct 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型" )
		default:
			fmt.Printf("未知型")
	}
	fmt.Println()
/*
select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的
1. 如果任意某个通信可以进行，它就执行；其他被忽略
2. 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行
3. 没有匹配时
	如果有default子句，则执行该语句
	如果没有default字句，select将阻塞，直到某个通信可以运行
*/
	chanCap := 5
	ch := make(chan int, chanCap) //创建channel，容量为5
	for i := 0; i < chanCap; i++ { //通过for循环，向channel里填满数据
		select { //通过select随机的向channel里追加数据
			case ch <- 1:
			case ch <- 2:
			case ch <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v ", <-ch)
	}//2 2 1 1 3
	fmt.Println()
/*******************************************************************循环语句*/
/*循环控制语句
break 语句, continue 语句, goto 语句(将控制转移到被标记的语句)
for _, num := range nums {sum += num} "_"是空白符,不需要索引时使用此符号
*/
	/* 定义局部变量 */
	var va0 int = 1
	LOOP: for va0 < 5 {
		if va0 == 3 {
			/* 跳过迭代 */
			va0 = va0 + 1
			goto LOOP
		}
		fmt.Printf("goto va0 的值为 : %d\n", va0)
		va0++
	}
	//goto va0 的值为 : 1
	//goto va0 的值为 : 2
	//goto va0 的值为 : 4
	var va1 int = 6
	var va2 int
	for va2 := 0; va2 < 3; va2++ {
		fmt.Printf("第一个for va2 的值为: %d\n", va2)
	}
	//第一个for va2 的值为: 1
	//第一个for va2 的值为: 2
	for va2 < va1 {
		va2++
		fmt.Printf("第二个for va2 的值为: %d\n", va2)
	}
	//第二个for va2 的值为: 1
	//第二个for va2 的值为: 2
	//第二个for va2 的值为: 3
	//第二个for va2 的值为: 4
	//第二个for va2 的值为: 5
	//第二个for va2 的值为: 6
	numbers := [6]int{2, 6, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	//第 0 位 x 的值 = 2
	//第 1 位 x 的值 = 6
	//第 2 位 x 的值 = 5
	//第 3 位 x 的值 = 0
	//第 4 位 x 的值 = 0
	//第 5 位 x 的值 = 0
	//for true  {
	//	fmt.Printf("这是无限循环。\n");
	//}
	//这是无限循环。
/*******************************************************************函数*/
/*
不同类型的局部和全局变量默认值为：
数据类型	初始化默认值
int			0
float32		0
pointer		nil
*/
/*函数参数有两种: 值传递值(实际参数复制一份传递), 引用传递(地址传递) */
	/*func max(num1, num2 int) int {
		var result int
		if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
		return result
	}
	fmt.Printf( "最大值是 : %d\n", max(100, 200))*/
	//最大值是 : 200

	/*返回多个值
	func swap(x, y string) (string, string) {
		return y, x
	}
	vf1, vf2 := swap("Mahesh", "Kumar")
	fmt.Println(vf1, vf2)//Kumar Mahesh
	*/
	/*
	引用传递值
	func swap(x *int, y *int) {
		var temp int
		temp = *x    //保存 x 地址上的值
		*x = *y      //将 y 值赋给 x
		*y = temp    //将 temp 值赋给 y
	}
	var vfa1 int = 100
	var vfa2 int = 200
	fmt.Printf("引用传递值前，vfa1 的值 : %d\n", vfa1 )
	fmt.Printf("引用传递值前，vfa2 的值 : %d\n", vfa2 )
	//引用传递值前，vfa1 的值 : 100
	//引用传递值前，vfa2 的值 : 200
	swap(&vfa1, &vfa2)
	fmt.Printf("引用传递值后，vfa1 的值 : %d\n", vfa1 )
	fmt.Printf("引用传递值后，vfa2 的值 : %d\n", vfa2 )
	//引用传递值后，vfa1 的值 : 200
	//引用传递值后，vfa2 的值 : 100
	*/

	/*func getAverage(arr []int, size int) float32 {
		var i,sum int
		var avg float32
		for i = 0; i < size;i++ {
		sum += arr[i]
	}
		avg = float32(sum) / float32(size)
		return avg;
	}
	var farr = []int {1000, 2, 3, 17, 50}
	var avg float32
	avg = getAverage( farr, 5)
	fmt.Printf( "平均值为: %f ", avg )//平均值为: 214.399994
	*/

	fmt.Println()
/*******************************************************************数组*/
/*
数组的长度不可改变
*/
	//一维数组的定义
	var arr1 [10] float32
	//初始化数组
	var arr2 = [5] float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//不设置数组大小
	var arr3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("一维数组:", arr1, arr2, arr3)//一维数组: [0 0 0 0 0 0 0 0 0 0] [1000 2 3.4 7 50] [1000 2 3.4 7 50]
	//二维的整型数组
	//var arr4 [3][4]int
	//初始化二维数组
	var arr5 = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	/* 输出数组元素 */
	for  i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, arr5[i][j] )
		}
	}
	//a[0][0] = 0
	//a[0][1] = 0
	//a[1][0] = 1
	//a[1][1] = 2
	//a[2][0] = 2
	//a[2][1] = 4
	//a[3][0] = 3
	//a[3][1] = 6
	//a[4][0] = 4
	//a[4][1] = 8
/*******************************************************************切片(Slice)*/
/*
功能强悍的内置类型切片("动态数组")
切片在未初始化之前默认为 nil，长度为 0
1. var slc1 [] int//未指定大小的数组来定义切片
   slc2 :=[] int {1, 2, 3}//切片初始化
2. var slice1 []type = make([]type, len)//len 是数组的长度并且也是切片的初始长度
   slice2 :=make([]int, len, cap)//cap指定容量
函数
len() 方法获取长度
cap() 可以测量切片最长可以达到多少
copy() 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
append() 向切片追加新元素

*/
	var slice1 = make([]int,3,5)
	fmt.Printf("切片(Slice) len=%d cap=%d slice=%v\n",len(slice1),cap(slice1), slice1)//切片(Slice) len=3 cap=5 slice=[0 0 0]
	slice2 := []int{0,1,2,3,4,5,6,7,8}
	fmt.Printf("切片(Slice) len=%d cap=%d slice=%v\n",len(slice2),cap(slice2), slice2)
		//切片(Slice) len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
	fmt.Println("切片(Slice) slice2 ==", slice2)
		//切片(Slice) slice2 == [0 1 2 3 4 5 6 7 8]
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("切片(Slice) slice2[1:4] ==", slice2[1:4])
		//切片(Slice) slice2[1:4] == [1 2 3]
	/* 默认下限为 0*/
	fmt.Println("切片(Slice) slice2[:3] ==", slice2[:3])
		//切片(Slice) slice2[:3] == [0 1 2]
	/* 默认上限为 len(s)*/
	fmt.Println("切片(Slice) slice2[4:] ==", slice2[4:])
		//切片(Slice) slice2[4:] == [4 5 6 7 8]

	var slice3 []int
	slice3 = append(slice3, 64)
	fmt.Printf("切片(Slice)3 len=%d cap=%d slice=%v\n",len(slice3),cap(slice3), slice3)
		//切片(Slice)3 len=1 cap=1 slice=[64]
	slice3 = append(slice3, 55,22,77)
	fmt.Printf("切片(Slice)3 len=%d cap=%d slice=%v\n",len(slice3),cap(slice3), slice3)
		//切片(Slice)3 len=4 cap=4 slice=[64 55 22 77]

	//创建切片 slice4 是之前切片的两倍容量
	slice4 := make([]int, len(slice3), (cap(slice3))*2)
	fmt.Printf("切片(Slice)4 len=%d cap=%d slice=%v\n",len(slice4),cap(slice4), slice4)
		//切片(Slice)4 len=4 cap=8 slice=[0 0 0 0]

	//拷贝 slice3 的内容到 slice4
	copy(slice4, slice3)
	fmt.Printf("切片(Slice)4 len=%d cap=%d slice=%v\n",len(slice4),cap(slice4), slice4)
	//切片(Slice)4 len=4 cap=8 slice=[64 55 22 77]

/*******************************************************************指针*/
/*
变量是一种使用方便的占位符，用于引用计算机内存地址,
取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
var var_name *var-type: var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针
指针类型前面加上 * 号（前缀）来获取指针所指向的内容
*/
	var v_pt1 *int        /* 指向整型*/
	var v_pt2 *float32    /* 指向浮点型 */
	fmt.Println("类型指针", v_pt1, v_pt2)//类型指针 <nil> <nil>
	fmt.Println("类型指针", &v_pt1, v_pt2)//类型指针 0xc042074020 <nil>
	fmt.Println()

	var v_1 int = 20	/* 声明实际变量 */
	var v_pt3 *int    	/* 声明指针变量 */
	v_pt3 = &v_1  		/* 指针变量的存储地址 */
	fmt.Printf("v_1 变量的地址是: %x\n", &v_1  )		//v_1 变量的地址是: c042054260
	/* 指针变量的存储地址 */
	fmt.Printf("v_pt3 变量储存的指针地址: %x\n", v_pt3 )//v_pt3 变量储存的指针地址: c042054260
	/* 使用指针访问值 */
	fmt.Printf("*v_pt3 变量的值: %d\n", *v_pt3 )		//*v_pt3 变量的值: 20

/*
指针数组
*/
	v_arr1 := []int{10, 100, 200}
	var ptr1 [3]*int
	var size = 3
	for  i := 0; i < size; i++ {
		ptr1[i] = &v_arr1[i] /* 整数地址赋值给指针数组 */
	}
	for  i := 0; i < size; i++ {
		fmt.Printf("arr1[%d] = %d\n", i, *ptr1[i])
	}
	//arr1[0] = 10
	//arr1[1] = 100
	//arr1[2] = 200

/*
指向指针的指针
声明格式:　var ptr **int
*/
	var v_2 int = 3000
	var ptr2 *int
	var pptr2 **int
	/* 指针 ptr2 地址 */
	ptr2 = &v_2
	/* 指向指针 ptr2 地址 */
	pptr2 = &ptr2
	/* 获取 pptr2 的值 */
	fmt.Printf("变量 v_2 = %d\n", v_2 )							//变量 v_2 = 3000
	fmt.Printf("指针变量 *ptr2 = %d\n", *ptr2 )					//指针变量 *ptr2 = 3000
	fmt.Printf("指向指针的指针变量 **pptr2 = %d\n", **pptr2)	//指向指针的指针变量 **pptr2 = 3000

/*
指针作为函数参数
*/
	/*func swap(x *int, y *int) {
		var temp int
		temp = *x    //保存 x 地址的值
		*x = *y      //将 y 赋值给 x
		*y = temp    //将 temp 赋值给 y
	}
	var v_3 int = 100
	var v_4 int = 200
	//调用函数用于交换值
	//&v_3 指向 v_3 变量的地址
	//&v_4 指向 v_4 变量的地址
	swap(&v_3, &v_4)
	fmt.Printf("交换后 v_3 的值 : %d\n", v_3 )//交换后 v_3 的值 : 200
	fmt.Printf("交换后 v_4 的值 : %d\n", v_4 )//交换后 v_4 的值 : 100
*/
/*******************************************************************结构体*/
	/*type Books struct {
		title string
		author string
		subject string
		book_id int
	}
	func printBook(book *Books) {//指针参数
		fmt.Printf( "Book ptr title : %s\n", book.title);
		fmt.Printf( "Book ptr author : %s\n", book.author);
		fmt.Printf( "Book ptr subject : %s\n", book.subject);
		fmt.Printf( "Book ptr book_id : %d\n", book.book_id);
	}
	var Book1 Books        //声明 Book1 为 Books 类型
	Book1.title = "标题"
	Book1.author = "www.shawn.com"
	Book1.subject = "主题"
	Book1.book_id = 108
	fmt.Printf( "Book 1 title : %s\n", Book1.title)//Book 1 title : 标题
	fmt.Printf( "Book 1 author : %s\n", Book1.author)//Book 1 author : www.shawn.com
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)//Book 1 subject : 主题
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)//Book 1 book_id : 108
	//传送结构体地址
	printBook(&Book1)*/
	//Book ptr title : 标题
	//Book ptr author : www.shawn.com
	//Book ptr subject : 主题
	//Book ptr book_id : 108

/*******************************************************************范围(Range)*/
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {//range去求一个slice的和, "_"是空白符,不需要索引时使用此符号
		sum += num
	}
	fmt.Println("sum:", sum)//sum: 9

	for i, num := range nums {//i为索引
		if num == 3 {
			fmt.Println("index:", i)//index: 1
		}
	}


	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {//range也可以用在map的键值对上
		fmt.Printf("%s -> %s\n", k, v)
	}
	//a -> apple
	//b -> banana

	for i, c := range "go" {	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身
		fmt.Println(i, c)
	}
	//0 103
	//1 111

/********************************************************************/




}
