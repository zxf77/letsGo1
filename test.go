//当前程序的包名
package main

//导入的包
import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//定义常量
const PI = 3.14159

//简写定义
const (
	test11 = 1
	test22 = "test22"
)

//常量的枚举，枚举存储单位
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	//这些都自动套用上一行的表达式
)

//定义全局变量并赋值
var name = "gotest"

//一般类型的声明
var newType int

//声明结构
type gopher struct {
}

//声明接口
type golang interface {
}

//可以给关键字声明别名
type (
	testdd int
	文本     string
)

//定义结构
type myStructPerson struct {
	Name string
	Age  int
}

//结构里面会有匿名结构
type myNoNameStruct struct {
	Name string
	Age  int
	//这下面是匿名结构
	Contact struct {
		Phone, City string
	}
}

//定义接口
type USB interface {
	Name() string //方法和返回类型是string
	Connect()
}
type PhoneConnector struct {
	name string
}

//实现接口 的方法
func (pc PhoneConnector) Name() string {
	return pc.name
}

//实现接口的另一个方法
func (pc PhoneConnector) Connect() {
	fmt.Println("通过接口实现的方法Connect:", pc.name)
}

//学习反射 定义一个结构
type User struct {
	Id   int
	Name string
	Age  int
}

//这个结构的方法
func (u User) Hello() {
	fmt.Println("结构的方法 hello")
}

//main()函数是程序入口点
func main() {
	fmt.Println("hello wordl!")
	fmt.Println("测试一下")

	//测试一下可见性规则
	fmt.Println()

	//验证一下各个类型的零值
	var a 文本
	a = "中文名字"

	fmt.Println("类型的0值", a)

	/*var testint int	//古老的声明变量， 当前不赋值，后续再使用的时候赋值
	testint = 1 	//古老的变量赋值*/

	//var testtest = 33 	//声明的同时直接赋值  类型由系统推断
	test11 := 1.3 //最简单的写法
	fmt.Println("变量声明", test11)

	var a1 int = 65
	b := strconv.Itoa(a1)
	fmt.Println(b)

	//左移运算符的使用,枚举存储单位
	/**
	1: 0001
	*/
	fmt.Println("左移运算符，存储单位", MB)

	//指针学习
	a2 := 1
	var p *int = &a2
	fmt.Println("指针", *p)

	//递增递减语句，不能作为表达式，只能是语句
	a2++

	//条件表达式
	//可以初始化表达式，且局部变量
	if a := 2; a > 1 {
		fmt.Println("初始化")
	}
	//循环语句for
	for {
		a2++
		if a2 > 5 {
			break
		}
		fmt.Println("循环的语句", a2)
	}
	//第二种形式
	for a2 < 8 {
		a2++
		fmt.Println("第二种形式", a2)
	}
	//第三种经典形式
	for i := 0; i < 3; i++ {
		fmt.Println("经典形式", i)
	}

	//选择语句switch
	bb := 3
	//第一种形式
	switch bb {
	case 0:
		fmt.Println("选择", "a=0")
	case 1:
		fmt.Println("选择", "a=1")
	default:
		fmt.Println("None")
	}
	//第二种形式
	switch {
	case bb >= 0:
		fmt.Println("switch第二个形式", "a >= 0")
		fallthrough //这个会判断下面的case
	case bb >= 1:
		fmt.Println("switch第二个形式", "a >= 1")
		fallthrough
	case bb >= 2:
		fmt.Println("switch第二个形式", "a >= 2")
	default:
		fmt.Println("none")
	}

	//跳转语句goto break continue
	//continue常用与退出有限循环中的无限循环
	//goto是调整执行位置
	//break是跳转到标签的同一层级别
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL1
			}
		}
	}
LABEL1: //标签
	fmt.Println("跳转到了这里")

	/**
	数组
	数组的长度也是数组类型的一部分
	*/
	//定义数组，可以通过冒号赋值给这个索引。数组长度可以是.....如果没有完整定义会尽量满足要求
	arr2 := [...]int{10: 1}
	arr := [3]int{1, 2: 3}
	fmt.Println("数组", arr, arr2)
	//区分数组的指针和指针数组
	var p2 *[3]int = &arr //数组的指针
	fmt.Println("数组的指针", p2)
	//指针数组
	a3, b3 := 1, 2
	arr3 := [...]*int{&a3, &b3}
	fmt.Println("指针数组", arr3)
	//使用new创建数组返回的是指向数组的指针
	ppp := new([5]int)
	fmt.Println(ppp)

	//多维数组..同样可以通过索引加冒号的值来定义某个索引的值
	arr4 := [2][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Println(arr4)

	/**
	go语言版的冒泡排序
	*/
	arrMaopao := [...]int{3, 4, 1, 6, 9}
	fmt.Println("冒泡排序前", arrMaopao)
	num := len(arrMaopao) //因为go语言的for循环每次遍历都会比较一次。所以提前计算一下变量
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arrMaopao[i] > arrMaopao[j] {
				temp := arrMaopao[i]
				arrMaopao[i] = arrMaopao[j]
				arrMaopao[j] = temp
			}
		}
	}
	fmt.Println("冒泡排序后", arrMaopao)

	//切片Slice
	arrS1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := arrS1[5:10] //取数组索引为5-9的元素，不包含索引10
	s1 = arrS1[:5]    //前五个元素,不包含索引5
	s1 = arrS1[5:]    //第五个到最后
	fmt.Println("打印切片", s1)

	//使用make定义切片
	s2 := make([]int, 3)
	fmt.Println("切片的信息", len(s2), cap(s2))
	fmt.Println("切片内容", s2)

	bb1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := bb1[2:5]
	fmt.Println(s3, len(s3), cap(s3))
	//个数就是3，容量则是从三开始到最后是7 - 2 + 1
	s3 = append(s3, 1, 2)
	fmt.Println("追加后", s3)
	//数组拷贝,把后一个数组copy给前一个数组。其中长度以短的为准
	sCopy1 := []int{1, 2, 3, 4, 5, 6}
	sCopy2 := []int{10, 12, 77}
	copy(sCopy2, sCopy1)
	fmt.Println("copy后的数组", sCopy2)

	//map
	//创建map
	m := make(map[int]string)
	m[1] = "ok11" //赋值
	delete(m, 1)  //删除键值对
	fmt.Println("简单map", m[1])
	//map嵌套map需要注意每一级的map都需要单独的初始化
	var mm2 map[int]map[int]string
	mm2 = make(map[int]map[int]string)
	aa2, isOk := mm2[1][1] //第二层为bool代表是否存在这个键值对
	//fmt.Println(isOk)
	if !isOk {
		mm2[1] = make(map[int]string) //初始化第二层map
	}
	mm2[1][1] = "seven"
	aa2, isOk = mm2[1][1]
	fmt.Println("复杂map", aa2, isOk)
	//map的迭代操作 range..两参数i，j  如果迭代的是map，则第一个参数返回的是每一个map的key
	sm := make([]map[int]string, 5) //以map为类型的切片slice，五个元素
	for key := range sm {           //这里面的第二个参数只是拷贝，不会对原来的造成任何影响..第一个参数是下标
		/*v = make(map[int]string, 1)	//初始化map
		v[1] = "ok"
		fmt.Println(v)*/
		sm[key] = make(map[int]string, 1)
		sm[key][1] = "okk"
		fmt.Println(sm[key])
	}
	fmt.Println(sm)
	//map的间接排序
	mm3 := map[int]string{1: "aa", 2: "bbb", 7: "ccc", 4: "ddd", 5: "eee"}
	ss3 := make([]int, len(mm3))
	ii3 := 0
	for key, _ := range mm3 {
		ss3[ii3] = key
		ii3++
	}
	sort.Ints(ss3)
	fmt.Println("go是无序的", ss3)
	/**作业：
	将map[int]string类型的一组map 键和值交换，变成类型map[string]int
	*/
	map1Homework := map[int]string{1: "aa", 2: "bbb", 7: "ccc", 4: "ddd", 5: "eee"}
	map2HomeWork := make(map[string]int)
	for key, val := range map1Homework {
		map2HomeWork[val] = key
	}
	fmt.Println("转换类型后", map2HomeWork)
	//myFunc3(1, 2, 3, 4, 5)
	//匿名函数,把这个函数作为一个类型传递给变量
	aaFunc := func() {
		fmt.Println("这是匿名函数")
	}
	aaFunc()
	//闭包的使用
	funcClose1 := clouure(3)           //先定义一个变量接受闭包函数返回的匿名函数
	fmt.Println("使用闭包", funcClose1(4)) //使用闭包

	/**
	defer
	在函数体执行结束后，按着调用顺序的 **相反顺序** 逐个执行
	*/
	/*fmt.Println("以下是defer，逆序向上调用", "a")
	defer fmt.Println("b")
	defer fmt.Println("c")*/
	//for循环来使用defer
	for i := 0; i < 3; i++ {
		//defer fmt.Println(i)
	}
	//defer的高级调用

	//结构
	structTest1 := myStructPerson{
		//字面值初始化的方法
		Name: "joe",
		Age:  17,
	}
	structTest1.methodTest()
	/*一般方法初始化
	structTest1.Name = "zhengxf"
	structTest1.Age = 19*/
	fmt.Println("自己的结构体struct", structTest1)
	//结构初始化建议使用取地址符号,这样修改或者调用函数可以直接引用
	structTest2 := &myStructPerson{
		Name: "zhengxf",
		Age:  23,
	}
	fmt.Println("指针的取地址,函数修改前", structTest2)
	changeStruct(structTest2)
	fmt.Println("函数修改后", structTest2)
	//匿名结构.同样可以加上取地址符来变成引用类型。
	noNameStruct := struct {
		Name string
		Age  int
	}{
		Name: "zzzz",
		Age:  111,
	}
	fmt.Println("匿名结构", noNameStruct)
	//结构里面有匿名结构
	noNameStruct1 := myNoNameStruct{
		Name: "test",
		Age:  22,
	}
	noNameStruct1.Contact.Phone = "1313131"
	noNameStruct1.Contact.City = "bj"
	fmt.Println("结构里面有匿名结构", noNameStruct1)
	//字符串String
	//直接使用+ 拼接字符串
	str2 := "测试拼接字符串hello" + " " + "world"
	//简写的+=形式
	str2 += "拼接成功"
	fmt.Println(str2)
	//其它方法拼接字符串
	str3 := fmt.Sprintf("%d : %s", 2021, "年拼接字符串") //逻辑复杂 用了interface 性能一般般
	fmt.Println(str3)
	fmt.Println(strings.Join([]string{"hello", "world测试拼接"}, ",")) //这个方法已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小。
	//还可以使用bytes.Buffer  strings.Builder(非线程安全)

	//方法method
	var methodTest USB //接口
	methodTest = PhoneConnector{"Phone"}
	methodTest.Connect()

	//反射
	uu := User{1, "ok", 22}
	info(uu)

	/**
	并发学习，
	*/
	//goroutine最简单的用法
	/*go Go()
	time.Sleep(2 * time.Second)	//休眠两秒钟*/
	//通过channel来通信
	ccc := make(chan bool)
	go func() {
		fmt.Println("GO通过通信来并发")
		ccc <- true

		//最好需要手动关闭
		close(ccc)
	}()
	for val := range ccc { //可以forr一直迭代channel
		fmt.Println("到这了", val)
	}

}
func test() {
	const Name = "123"
}

//定义自己的函数function
//两个参数，返回一个int
func myFunc(a int, b string) int {
	return 0
}

//参数a和b都是int类型，返回类型是int和string两个
func myFunc1(a, b int) (int, string) {
	return int(1), string("test")
}

//命名返回值，若使用了命名返回值，则return就可以
func myFunc2(a int) (d, b, c int) {
	return d, b, c
}

//go语言的不定长变参
func myFunc3(a ...int) {
	fmt.Println("调用了函数", a)
}

//闭包函数，作用是返回一个匿名函数
func clouure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

/**
修改结构内容的方法，传参用指针，声明结构的时候加上&。传参的时候不用&
*/
func changeStruct(person *myStructPerson) {
	person.Age = 77
	fmt.Println("函数中修改结构", person)
}

/**
方法 method  通过接收者receiver判断是哪个结构的方法,第一个括号定义的哪个结构
如果方法的结构不同，则可以重名，通过结构体.method 来判断调用哪个方法
receiver也同样可以用值传递或者引用指针传递
方法可以访问私有字段
*/
func (a myStructPerson) methodTest() {
	fmt.Println("这是方法method")
}

/**
反射使用的函数
参数是一个空接口
*/
func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("反射获取到的接口类型type:", t.Name(), t)

	v := reflect.ValueOf(o)
	fmt.Println("Fields:", v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}

/**
并行学习定义函数
*/
func Go() {
	fmt.Println("gogogo 并行")
}
