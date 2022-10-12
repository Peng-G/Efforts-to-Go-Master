package go_basics

import "fmt"

func Arr_and_slice() {
	a := [3]int{1, 2, 3}  // 1. 数组，长度固定，类型相同，声明时必须指定长度，长度不可变，数组是值类型，赋值和传参时会复制整个数组，而不是指针，因此改变副本的值，不会改变本身的值，数组之间不可比较，只能和nil比较，数组可以通过new函数分配内存，但是不会初始化元素，数组的每个元素都是零值，数组的长度是数组类型的一部分，因此[3]int和[4]int是不同的类型。
	fmt.Printf("Grads: %v\n", a)

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}  // 2. 数组的长度是数组类型的一部分，可以省略长度，由编译器自动推断，但是不建议这么做，因为这样会影响代码的可读性，而且如果数组的长度变化了，编译器不会报错，但是程序运行时会出错。
	fmt.Printf("Arr: %v\n", arr)

	s := []int{1, 2, 3}  // 3. 切片，与Python中的列表类似，长度可变，类型相同，声明时不需要指定长度，切片是引用类型，底层是数组，切片的长度可以改变，但是不能超过底层数组的容量，切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数，切片的零值是nil，切片的长度和容量可以通过内置的len()和cap()函数来获取。
	fmt.Printf("S[0]: %v\n", s[0])
	fmt.Printf("S[1:]: %v\n", s[1:])

	matrix := [3][3]int{  // 4. 多维数组，数组的元素也可以是数组，多维数组的长度是数组类型的一部分，因此[3][3]int和[4][3]int是不同的类型。
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9},
	}
	fmt.Printf("Matrix: %v\n", matrix)

	b := [3]int{1, 2, 3}  // 5. 数组是值类型，赋值和传参会复制整个数组，而不是指针，因此改变副本的值，不会改变本身的值。
	c := b
	c[0] = 100
	fmt.Printf("B: %v, C: %v\n", b, c)

	d := &b  // 6. 使用数组的指针，可以避免数组的复制，改变指针指向的数组的值，也会改变本身的值。
	d[0] = 100
	fmt.Printf("B: %v, D: %v\n", b, d)

	e := []int{1, 2, 3}  // 7. 切片是引用类型，赋值和传参会复制指针，指向同一个底层数组。
	f := e
	f[0] = 100
	fmt.Printf("E: %v, F: %v\n", e, f)

	g := make([]int, 3, 10)  // 8. 使用make函数创建切片, make([]T, len, cap)，其中T是切片的元素类型，len是切片的长度，cap是切片的容量，cap可以省略，此时cap=len。
	fmt.Printf("G: %v\n", g)  // 
	fmt.Printf("Length: %v, Capacity: %v\n", len(g), cap(g)) // 3, 10

	h := []int{}
	fmt.Printf("H: %v, Length: %v, Capacity: %v\n", h, len(h), cap(h))  // 9. 空切片的长度和容量都是0
	h = append(h, 1)  // 10. 使用append函数添加元素, append(slice, elements...)，其中slice是原切片，elements是要添加的元素，可以添加一个或多个元素，返回值是一个包含原切片所有元素加上新添加元素的切片。
	fmt.Printf("H: %v, Length: %v, Capacity: %v\n", h, len(h), cap(h))  // 1, 1, 1
	h = append(h, 2, 3, 4) // 11. append函数可以一次添加多个元素，如果原切片的容量不够，append函数会分配一个更大的底层数组，然后将原数组的内容复制到新数组中，返回的切片指向新数组。
	fmt.Printf("H: %v, Length: %v, Capacity: %v\n", h, len(h), cap(h))  // 1, 2, 3, 4, 4, 4

	// 12. 切片的容量是指底层数组的容量，切片的长度是指切片中元素的个数，如果可以预知切片的长度，可以使用make函数创建切片，避免切片的扩容，因为扩容会导致底层数组的复制，影响性能。

	i := []int{1, 2, 3, 4, 5}
	j := i[:len(i)-1]  // 13. 切片的切片，可以使用切片表达式创建切片的切片，切片表达式的语法是slice[low:high]，其中low和high都是可选的，low默认为0，high默认为len(slice)，返回值是一个新的切片，指向原切片的底层数组。
	fmt.Println(j)  // [1 2 3 4]
	k := append(i[:2], i[3:]...)  // 14. 删除切片中的元素，使用append函数，将切片的前半部分和后半部分拼接起来，注意...的使用，表示将切片展开，作为append函数的参数，而不是作为一个元素。
	fmt.Println(k)  // [1 2 4 5]
	fmt.Println(i)  // [1 2 4 5 5] // 15. 切片是引用类型，删除元素后，原切片也会改变，因为切片的底层数组是共享的。
}