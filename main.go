package main

import (
	"fmt"
	"math/rand"
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
	var a int8
	b := uint8(a)
}

// 格式化打印
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
}

// 字符串操作
func handle_string() {
	var str string = "abc"
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

	// char始终为对应集合中的值拷贝
	// 如果是字符串，可以自动按utf8识别
	for pos, char := range str {

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

// 数组
func array() {
	
}

func init() {
	fmt.Println("在初始化包的时候执行，包的每个源文件都可以定义init函数")
}

func main() {
	//format_print()
}
