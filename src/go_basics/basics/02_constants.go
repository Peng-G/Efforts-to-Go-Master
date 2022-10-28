package go_basics

import (
	"fmt"
)
// 4. 常量的定义可以放在函数外面
const (
	a = iota // 5. iota是Go语言的常量计数器，只能在常量的表达式中使用
	b = iota // 6. iota在const关键字出现时将被重置为0，const中每新增一行常量声明将使iota计数一次
	c = iota // 7. iota可以被用作枚举值
) 

const (
	_ = iota// 8. iota在每遇到一个const关键字时会重置，所以每遇到一个const关键字就相当于重新开始声明常量
	d 
	e
	f
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (  // 9. 常量的定义可以分组
	isAdmin = 1 << iota  // 10. 可以使用常量来定义一组布尔值
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func Constants() {
	const MY_CONST int = 100  // 1. 常量的定义
	const myConst = 200  // 2. 常量的类型可以省略，根据值自行判断
	fmt.Println(MY_CONST, myConst)

	var v int16 = 200  // 3. 常量可以和其他类型的变量进行运算
	fmt.Printf("%v, %T\n", myConst + v, myConst + v) // 400, int16

	fmt.Println(a, b, c, d, e, f) // 0 1 2 0 1 2
	
	fmt.Println(KB, MB, GB, TB) // 1024 1048576 1073741824 1099511627776
	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB) // 3.73GB

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles) // 1010001
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin) // Is Admin? true
	fmt.Printf("Is HQ? %v\n", isHeadquarters & roles == isHeadquarters) // Is HQ? false

}
