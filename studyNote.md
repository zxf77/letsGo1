## go学习

Go并发支持、垃圾回收的编译型系统编程语言，旨在创造一门具有静态编译语言的高性能和动态语言的高效开发之间的良好平衡。

Go的主要特点：

* 类型安全 和 内存安全
* 以直观和极低的代价实现高并发
* 高效的垃圾回收机制
* 快速编译（解决c语言头文件太多）
* 为多核计算机提供性能提升的方案
* UTF—8编码支持



#### go程序的一般结构

* go程序通过package来组织
* 只有package名称为main的包可以含有main函数
* **一个可执行程序 有且仅有一个main包**
* import导入其他飞main包
* 通过const关键字定义常量
* 函数体外使用var关键字 申明和复制全局变量
* type关键字声明结构（struct）和借口（interface）
* 通过func声明函数

go语言的可见性规则：

使用大小写来决定 常量、变量、类型、接口、结构或函数，根据约定，函数名首字母小写即为private，大写即为public。



### go基本类型

* 特殊的：足够保存执政的32位或64位整数幸：uintptr
* 其它值类型：array、struct、string
* 引用类型：slice（切面）、map、chan（通道）
* 借口类型interface。函数类型func

#### 类型零值

零值不等于空值。指默认值，通常下默认值是0，bool为false。

#### 多个变量的声明和赋值

------

1. 全局变量声明可以使用var()缩写
2. 全局变量不可以参略var，但可以并行
3. 所有变量都可以使用类型推断
4. 局部变量不可以用时var（）的方式简写，只能使用并行方式，可以省略var关键字

#### 变量的类型转换

go不存在隐式转换，all类型必须显示声明

转换只能发生在两种**相互兼容的类型之间**

```
strconv.Itoa(a1)	//int转换成string保持数字
```



#### 常量

* 定义时候若不提供初始值，则将使用上行表达式的值
* 相同的表达式不代表具有相同的值
* iota常量计数器，从0开始，没定一个常量自动递增1
* 通过初试规则和iota可以达到枚举的效果
* 没遇到一个const关键字，iota就会重置为0
* 常量一般全部大写。

#### 运算符

专门用于channle的运算符 <-

左移运算符：<<



#### go的指针

go保留了指针，go不支持指针运算和“->”运算符，直接采用“.”选择符来操作指针目标对象的成员

-----操作符“&”取变量地址，使用“*”来通过指针间接访问目标对象

-----指针的默认值是nil

递增递减语句：go中。++和--作为语句不是表达式。（语句的话就不能放在等号右边）

#### 控制语句

##### if判断

-----条件表达式没有括号

------支撑初始化表达式（可以使并行方式）

------左大括号必须和条件语句或else再同一行

------ 支持单行

------初始化语句变量为block，隐藏外部同名变量

##### 循环语句for

go只有一个for一个循环语句关键字，但支持三种形式

初始化和步进表达式可以是多个值

每次循环都会被重新检查，不建议再条件语句中使用函数，建议提前计算好条件并以变量或常量代替

左大括号必须和条件语句在同一行

##### 选择语句switch

可以使用任何类型或表达式

不需要写break，条件符合自动终止

若想执行下一个case，需要使用fallthrough

支持初始化表达式，右侧分号。初始化的变量同if语句一样 都是只在这个语句块中有效

左大括号和条件语句在同一行



##### 跳转语句goto，break ，continue

---三个语法都可以配合标签使用

---标签区分大小写，若不使用会造成编译错误

---break与continue配合标签可用于多层循环的跳出

--goto是调整执行位置，与其它两个语句配合标签的结果并不相同。



### 数组Array

1. 定义数组格式：var <varName> [n]<type>, n >= 0
2. **数组长度也是类型的一部分**，具有不同长度的数组为不同类型的数组
3. 需要区分数组的指针 和指针数组
4. go中数值为值类型
5. 数值可以使用==或！=比较，但不可以使用<或>
6. 可以使用new创建数组，此方法放回一个指向数组的指针
7. go支持多维数组

### 切片Slice

* 本身并不是数组，指向底层的数组
* 作为变长数组的替代方案，可以关联底层数组的局部或全部
* 为应用类型
* 可以直接创建或从底层数组获取生成
* 使用len()获取元素个数,cap()获取容量
* 一般使用make()创建
* **如果多个slice指向相同底层数组，其中一个值的改变会影响全部**
* 创建方式：make([]T, len, cap)，其中cap若省略则和len相同，len是个数，cap是容量
* 切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个 长度可变的数组。
* 绝对不要用指针指向 slice，切片本身已经是一个引用类型，所以它本身就是一个指针!

#### Reslice

---Reslice时索引以北slice的切片为准

--- 索引不可以超过北slice的切片的容量cap()值

--- 索引越界不会导致底层数组重新分配而是引发错误

#### Append

-- 可以在slice尾部追加元素

--可以将一个slice追加在另一个slice尾部

-- 若最终长度超过追加到slice的容量则返回原始slice

--若超过追加到的slice的容量则将重新分配数组并拷贝原始数据

#### copy

### map

* 类似哈希表或者字典，以key-value形式存储数据
* key必须是支持==或！=比较运算的类型，不可以是函数、map或slice
* map查找比线性搜索快。但比使用索引访问数据的类型慢很多
* map使用make()创建 支持 := 这种简写方式
* 定义方式：make([key Type]valueType, cap), cap表示容量，可省略
* 超出容量自动扩容、但尽量提供一个合理的初始值
* 使用len()获取元素个数
* 键值对不存在时自动添加，使用delete()删除某键值对
* 使用for range对map和slice进行迭代操作

### 函数function

1. go函数不支持嵌套、重载和默认参数
2. 但支持：无序声明原型，不定长度变参、多返回值、**命名返回值参数**、匿名函数、闭包
3. 定义函数使用func，左大括号不能另起一行
4. **函数也可以作为一种类型使用**
5. slice作为参数 传递的也是拷贝，不过传递的是地址的拷贝
6. 值类型作为参数是拷贝值。引用类型作为参数是拷贝地址

#### 匿名函数

把这个函数作为一个类型传递给变量

#### 闭包

作用是返回一个匿名函数



### defer

---执行方式，在函数体执行结束后，按着调用顺序的 **相反顺序** 逐个执行

--- 即使函数发生 严重错误 也会执行

--支持匿名函数的调用

-- 常用于资源清理、文件关闭、解锁以及记录时间等操作

通过于匿名函数配合可在return之后修改函数计算结果

--如果函数体内某变量作为defer是匿名函数的参数，则在定义defer时已经获得了拷贝，否则则是引用某个变量的地址

---go没有异常机制，担忧panic/recover模式来处理错误

panic可以在任何地方应发，但是**recover只有在defer调用的函数中有效**



### 结构struct

go中的struct和c中的struct非常相似，并且go没有class

使用type <Name> struct{}定义结构，名称遵循可见性规则

支持指向自身的指针类型成员

支持匿名结构，可用作成员或定义成员变量

匿名结构也可以用于map的值

可以使用字面值对结构进行初始化

允许直接通过指针来读写结构成员

想同类型的成员可进行直接拷贝赋值

支持 == 与 ！= 比较运算符，不支持 > 或<

支持匿名字段，本质上是定义了以某个类型名为名称的字段

嵌入结构作为匿名字段看起来像是继承，但不是继承

可以使用匿名字段指针

结构初始化的时候建议使用取地址符号，有利于后期的操作

### 字符串String

* go语言string类型是不可变
* string类型的零值是nil
* 可以直接使用 + 运算符拼接字符串，简写的+=形式也可以使用。不过会生成很多临时五用的字符串，性能较差
* 其它方法：fmt.Sprintf(),strings.Join()，还可以使用bytes.Buffer  strings.Builder(非线程安全)



### 方法method

* go没有class  依然有method
* 通过显示说明receiver来实现与某个类型的组合
* 只能为同一个包中的类型定义方法
* receiver可以是类型的值或者指针
* 不存在方法重载
* 可以使用值或指针来调用方法，编译器自动完成转换
* 某种意义上，方法是函数的语法糖，receiver其实就是方法所接收的第一个参数
* 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
* 类型别名不会拥有底层类型所附带的方法
* 方法可以调用结构中的私有字段



### 接口interface

1.   接口是一个或多个方法签名的集合
2. 只要某个类型有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个几口，称为Structural Typing
3. 接口只有方法声明，没有实现，没有数据字段
4. 接口可以匿名嵌入其他接口，或嵌入结构
5. 将对象赋值给接口，会拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品也无法获取指针
6. 只有当接口的存储类型和对象都为nil时，接口才等于nil
7. 接口调用不会做receiver的自动转换
8. 接口同样支持匿名字段方法
9. 接口可以实现类似oop的多态
10. 空接口可以作为任何类型数据的容器



### 反射reflection

反射可大大提高程序的灵活性，似的接口interface{}有更大的发挥余地

反射使用typeOf 和 ValueOf 函数从接口获取目标对象信息

反射 会将匿名字段作为独立字段（匿名字段本质）

想要利用反射修改对象的状态，前提是interface.data是settable,即pointer-interface

通过反射可以“动态”调用方法



### 并发concurrency

goroutine通过通信来共享内存 通过通道通信

#### channel通信

是goroutine沟通的桥梁，大都是阻塞同步

通过make创建， close关闭

Channle是引用类型

可以使用for range来迭代不断操作channel

可以设置单向或双向通道

可以设置缓存大小，有缓存则是异步的，在未被填满前不会发生阻塞。如果没缓存则是同步阻塞

#### Select

* 可以处理一个或多个channel的发送与接受
* 同时有多个可以用的channel时按随机顺序处理
* 可用空的select来阻塞main函数
* 可设置超时



### 项目相关

#### 完整的项目结构
