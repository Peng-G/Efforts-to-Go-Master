package go_basics

import (
	"fmt"
	// "io"
	// "log"
	// "net/http"
)

func Defer_panic_recover() {
	// 1. defer语句，延迟执行。
	defer fmt.Println("defer 1") // defer 语句会等到函数执行完毕后再执行。
	defer fmt.Println("defer 2") // defer 语句会按照先进后出的顺序执行。
	fmt.Println("main")

	// 2. defer语句会延迟执行，但是不会延迟函数的参数的计算。最后输出的是start，因为a的值在defer语句执行时已经确定了。
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// 3. panic语句，抛出异常。
	// panic("something bad happened")
	// 4. panic语句会中断函数的执行，但是会执行defer语句。

	// 5. recover语句，捕获异常，只在defer函数中有用，可以从panic中恢复函数。
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("something bad happened")

	// 6. 可以使用defer语句关闭资源，避免资源泄露，或者遗忘关闭资源。但是如果有太多资源被打开，defer语句会导致内存泄露。
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close() 
	// robots, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

}
