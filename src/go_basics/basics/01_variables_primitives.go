package go_basics

import "fmt"

// 2. 声明多个变量
var (
	name string = "Tom"  // 3. 声明的变量必须使用，否则会出现编译错误
	age  int    = 10
	isOk bool   = true  // 4. 声明变量时，可以省略类型，根据值自行判断变量类型
	// 5. 如果再声明变量时不赋值，会使用默认值0.
)
// 6. 命名规则：
//    1. 变量名必须以字母或下划线开头，不能以数字开头
//    2. 变量名不能包含空格
//    3. 不能使用Go的关键字和保留字作为变量名
//    4. 严格区分大小写
//    5. 小写的变量名只能在本包中使用，大写的变量名可以在本包和其他包中使用

func Variables_and_primitives() {
	// 1. 声明一个变量的三种方式
	var i int
	i = 1
	var j int = 2
	k := 3
	fmt.Println(i, j, k)
	fmt.Printf("%v, %T", k, k)

	// 7. 逻辑运算符
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1

	// 8. 位运算符
	c := 8
	fmt.Println(c << 3) // 64
	fmt.Println(c >> 3) // 1

	// 9. 复数
	var n complex64 = 1 + 2i  // 复数
	fmt.Printf("%v, %T\n", real(n), real(n)) // 1, float32
	fmt.Printf("%v, %T\n", imag(n), imag(n)) // 2, float32
	fmt.Printf("%v, %T\n", n, n) // (1+2i), complex64

	var m complex128 = complex(5, 12)  // 使用complex函数构造复数
	fmt.Printf("%v, %T\n", m, m) // (5+12i), complex128
	fmt.Printf("%v, %T\n", real(m), real(m)) // 5, float64
	fmt.Printf("%v, %T\n", imag(m), imag(m)) // 12, float64

	var s string = "hello"  // 字符串是不可变的
	fmt.Printf("%v, %T\n", s, s) // hello, string
	fmt.Printf("%v, %T\n", s[2], s[2]) // 108, uint8

	t := []byte(s)  // 将字符串强制转换为[]byte类型
	fmt.Printf("%v, %T\n", t, t) // [104 101 108 108 111], []uint8
	fmt.Printf("%v, %T\n", string(t), string(t)) // hello, string

	r := 'a'  // rune
	fmt.Printf("%v, %T\n", r, r) // 97, int32

}
