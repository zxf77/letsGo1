//当前程序的包名
package main

//导入的包
import "fmt"

//定义常量
const PI = 3.14159

//简写定义
const (
	test11 = 1
	test22 = "test22"
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

}
func test() {
	const Name = "123"
}
