//package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// 自定义类型
type TZ int
type binOp func(int, int) int

// 变量
func variate() {
	// 数据类型
	// bool
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64(尽可能地使用 float64，因为math包中所有有关数学运算的函数都会要求接收这个类型)
	// byte(uint8), rune(int32)
	// _(空白符)

	var a int = 0
	var b = 0
	var c, d = 0, 0
	e, f := 0, 0
	// 枚举
	var (
		g = 0
		h
		i = "this is string"
	)
	const j int = 0
	const k = 0
	// 常量枚举(没有赋值的常量会应用上一行的赋值表达式)
	const (
		l = 0
		m
	)

	//值类型
	//	int, float, bool, string, 数组, struct, 指针
	//引用类型
	// slice, map, channel interface
}

// 类型转换
func convert_type() {
	// int8 -> uint8
	var a int8
	b := uint8(a)

	// string <-> slice
	var str string = "abc"
	var s1 []byte = []byte(str)
	var s2 []byte; copy(s2, str)
	str = string(s1)
}

// 打印
func format_print() {
	a := 0
	fmt.Printf("整数 %d\n", 1)
	fmt.Printf("十六进制 %x或%X\n", 0x01)
	fmt.Printf("格式化浮点数 %g, 普通浮点数 %f, 科学计数浮点数 %e\n", 0.1, 0.1, 0.1)
	fmt.Printf("定长整数 %02d\n", 1)
	fmt.Printf("定长小数 %2.3g, g可换成f和e\n", 0.1)
	fmt.Printf("二进制 %b\n", 5)
	fmt.Printf("字符 %c\n", 'a')
	fmt.Printf("用整数表示字符 %v或%d\n", 'a')
	fmt.Printf("utf8格式字符 %U\n", 'a')
	fmt.Printf("地址 %p", &a)
	fmt.Printf("%v", myClass{a:1, b:2})
}

// 字符串
func handle_string() {
	// 在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数
	var str string = "abc"
	str2 := str[1:2]
	str2[0] = 'a'	// 替换字符串中的字符是非法的
	fmt.Print(len(str))
	fmt.Println(str[0], str[len(str)-1])
	// 在循环中使用+号并不是最高效的做法，而是使用strings.Join()，或者bytes.Buffer
	str = "hello" + "word"
	str += "!"
}

// 指针操作
func handle_pointer() {
	a := 0
	var p *int = &a
	fmt.Print(p == nil)
	// 不能获得常量地址
	const b = 0
	p = &b
}

// 随机数
func random() {
	rand.Seed(int64(time.Now().Nanosecond()))
	a := rand.Int()
}

// 流程控制
func flow_control() {
	str := "abc"

	if a := 0; a == 0 {
		goto LABEL1
	} else if a == 1 {
		goto LABEL2
	} else {

	}

	switch a := 0; a {
	case 0:
		// do nothing
	case 1: fallthrough
	case 2:
		fmt.Print("a")
	default:
	}

	switch a := 0; {
	case a < 0:
	case a > 0:
	default:
	}

	for a := 0; a < 10; a++ {
		continue
	}

	for a := 0; ;a++ {
		break
	}

	for {
		break
	}

	for ;; {
		break
	}

	// value始终为对应集合中的"值拷贝"
	// 如果是字符串，可以自动按utf8识别
	for pos, value := range str {

	}
	for pos := range str {

	}

	// 标签和goto之间不能出现新变量定义
LABEL1:
	fmt.Print("a")
LABEL2:
	fmt.Print("b")

	defer func() {
		fmt.Print("call after calling flow_control")
	}()
}

// 函数
func function() {
	// 函数参数自动根据形参是值类型和引用类型决定传值还是引用
	// 尽量使用带命名的返回值

	// go语言的所有内置函数
	// close, len, cap, new, make, copy, append
	// panic, recover, print, println
	// complex, real, imag

	// 变参函数
	a := func(args ...interface{}) (c, d int) {
		for _, v := range args {
			switch t := v.(type) {
			case int:
			case float32:
			case string:
			default:
			}
		}
	}

	// 带函数参数的函数
	b := func(a int, f func(int, int)(int, int)) {
		f(0, 0)
	}

	return
}

// 数组(值类型)
func array() {
	var arr1 [5]int = [5]int{1, 2, 4}
	var arr2 [5]int = [...]int{1, 2, 3, 4, 5}
	var arr3 [5]byte = [5]byte{3:'a', 4:'c'}
	arr1[0] = 0
	arr1[len(arr1)-1] = 4

	// 数组是一种值类型，在做赋值操作的时候，会全量拷贝
	// 将数组传递给函数，用指针或者切片来防止全量拷贝
	arr2 = arr1
}

// 切片(引用类型)
func slice() {
	// 切片在内存中的组织方式实际上是一个有 3 个域的结构体
	// {
	//	指向相关数组的指针
	//	切片长度
	//	切片容量
	// }
	var arr1 [5]int
	var s1 []int = arr1[2:3]
	var s2 []int = arr1[:]
	s3 := s2[:]		// s3和s2指向相同的数组
	s4 := []int{1, 2, 3, 4}
	s5 := make([]int, 0)	// 指向空数组的切片(数组存在)
	s6 := new([]int)		// 指针为空的切片(数组不存在)
}

// map下的元素不是变量，因此不能寻址,比如数组或者结构体
// 所以如果要修改map中的变量的值，需要将其变成指针形式
func Map() {
	var m1 map[int]string
	m2 := make(map[int]string)
	m3 := map[int]string{1:"he", 2:"be"}

	if _, ok := m1[0]; ok {

	}
}

func new_and_make() {
	// 都在堆上分配内存
	// make适用于slice, map, channel
	// 其余适用new
}

// 。。。用法
func dot_dot_dot() {
	// 提供变参函数
	var f1 = func (args ...string) {}
	// 将切片元素分散传递
	s1 := []int{}
	s2 := []int{}
	s1 = append(s1, s2...)
}

// ----------- struct ---------------
type parent struct {
	c int
}
type myClass struct {
	a int
	b int
	int // 匿名字段
	parent // 匿名字段
}
func (a *myClass) Area() float32 {
	return 0
}
func (a *myClass) Radius() float32 {
	return 0
}
func (a *myClass) String() {
	// fmt.Print和fmt.Println会自动调用类型的String()方法
	// 所以不要再String方法内部调用fmt的两个方法，会导致递归调用
}

// 结构体(值类型)
func Struct() {
	c1 := myClass{1, 1, 1, parent{1}}
	c2 := &myClass{a:1}
	c1.int = 1
	c1.c = 1
}
// ----------- struct ---------------

// ---------- interface -------------
type Shaper interface {
	Area() float32
}
type Circle interface {
	Shaper
	Radius() float32
}
var Any interface{} // 空接口
func Interface() {
	var c myClass
	var s Circle = &c
	s.Area()

	// 类型判断
	if v, ok := s.(*myClass); ok {
		v.Radius()
	}
	switch t := s.(type) {
	case *myClass:
		t.Radius()
	default:
		fmt.Println("error")
	}
	// 接口类型转换，测试一个接口是否实现了另一个接口
	// 只要实际对象实现了两个接口，那么由他转换的接口变量可以对另一个接口变量赋值
	if v, ok := s.(Shaper); ok {}
}
// ---------- interface -------------

// ----------- reflect -------------
func Reflect() {
	var x float64 = 0
	reflect.TypeOf(x) // float64
	reflect.ValueOf(x) // 0
	// 通过反射修改值需要注意
	// ValueOf会拷贝创建返回值，改变返回值不会影响原值
	// 需要传递&地址给ValueOf
	// 另外可以通过CanSet判断是否可修改
	a := reflect.ValueOf(&x)
	if a.CanSet() {
		a = a.Elem()
		a.SetFloat(1)
	}
}
// ----------- reflect -------------

// ------- goroutine,channel ----------
func goroutine_channel() {
	// make(chann, int)和缓冲通道make(chann, int, >=1)
	// 是完全不同两种模式，一个是同步，一个是异步

	var ch chan int = make(chan int)
	ch <- 1
	a := <- ch

	// 函数通道
	var chf chan func()

	// 关闭通道后，无法继续写入
	close(ch)
	// 检测通道是否关闭
	if v, ok := <- ch; ok {}
	// 自动检测是否关闭
	for v := range ch {}

	ch2 := make(chan string)

	// 选择器
	select {
	case a := <- ch:
		fmt.Println(a)
	case b := <- ch2:
		fmt.Println(b)
	default:
		fmt.Println("...")
	}

	// 计时器
	ticker := time.NewTicker(time.Second)
	for v := range ticker.C {}

	// 结束协程
	runtime.Goexit()
}
// ------- goroutine,channel ----------

func init() {
	fmt.Println("在初始化包的时候执行，包的每个源文件都可以定义init函数，它们无序执行")
}

func main() {
	// 在 gc 编译器下（6g 或者 8g）你必须设置 GOMAXPROCS 为一个大于默认值 1 的数值来允许运行时支持使用多于 1 个的操作系统线程
	// 否则所有协程都由单线程来承载
	// 对于 n 个核心的情况设置 GOMAXPROCS 为 n-1 以获得最佳性能
	runtime.GOMAXPROCS(2)
}

// 可寻址
// 不可寻址（不可变的、临时结果和不安全的）
//		常量的值。
//		基本类型值的字面量。
//		算术操作的结果值。
//		对各种字面量的索引表达式和切片表达式的结果值:
//		不过有一个例外，对切片字面量的索引结果值却是可寻址的。
//		对字符串变量的索引表达式和切片表达式的结果值。
//		对字典变量的索引表达式的结果值。
//		函数字面量和方法字面量，以及对它们的调用表达式的结果值。
//		结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
//		类型转换表达式的结果值。
//		类型断言表达式的结果值。
//		接收表达式的结果值。


// 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针
// 对字符串使用a += b操作会导致大量内存拷贝，因为字符串是不可变类型
// defer只有在函数返回时才执行
// 不要将一个指向切片的指针传递给函数
// 使用值类型时误用指针，因为编译器为认为需要创建一个对象，并将对象移动到堆上
// 尽可能的使用:=声明定义变量
// 尽可能用字符代替字符串
// 尽可能用切片代替数组
// 如果只想获取切片中的值，不需要值的索引，尽可能用for range去遍历
// 当数组元素是稀疏的，使用映射会降低内存消耗
// 初始化映射时指定容量
// 定义方法时，使用指针类型作为方法接受者