package go_basics

import (
	"fmt"
)

func Loops() {
	// 1. for循环。
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2. for循环，省略初始语句和后置语句。
	i := 0
	for i < 5 {  // 类似于while循环。
		fmt.Println(i)
		i++
	}

	// 3. for循环，省略初始语句、后置语句和条件表达式。
	i = 0
	for {
		if i >= 5 {
			break
		}
		if i == 3 {
			i++
			continue  // 跳过本次循环。
		}
		fmt.Println(i)
		i++
	}

	// 4. 使用两个变量
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// 5. 双层for循环，使用label跳出循环。
Loop:  // label
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}

	// 6. for循环，遍历数组，类似于python的for in循环。
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 7. for循环，遍历map。
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 8. for循环，遍历字符串。
	s := "hello"
	for i, v := range s {
		fmt.Println(i, v)
	}
}
