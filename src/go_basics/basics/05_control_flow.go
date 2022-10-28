package go_basics

import (
	"fmt"
)

func Control_flow() {
	// 1. if语句。
	if true {  
		fmt.Println("true")
	}

	// 2. if语句中声明变量。
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Pennsylvania": 12802503,
		"Illinois":   12801539,
		"Ohio":       11614373,
	}
	if pop, ok := statePopulations["California"]; ok {
		fmt.Println(pop)
	}

	// 3. if-else-if-else语句。
	number := 50
	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("the guess must be between 1 and 100")
	} else if guess < number {
		fmt.Println("too low")
	} else {
		fmt.Println("too high")
	}

	// 4. switch语句。
	switch i := 2 + 3; i {
	case 1, 5, 10: // 5. switch语句中的case可以有多个值。
		fmt.Println("one, five or ten")
	case 2, 3, 4:
		fmt.Println("two, three or four")
	default: // 6. switch语句中的default分支是可选的。
		fmt.Println("something else")
	}

	j := 10
	switch { // 7. switch语句中可以不带表达式。
	case j <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // 8. fallthrough语句会强制执行后面的case语句，不管条件是否满足。
	case j <= 20:  // 9. switch语句是自动break的，除非使用fallthrough关键字。
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	var k interface{} = 1  // 10. switch语句中的类型开关。
	switch k.(type) { // 11. switch语句中可以使用类型断言。
	case int:  // 12. 冒号后面的所有语句都会被执行，直到遇到break语句，或者switch语句结束。
		fmt.Println("int")
		fmt.Println("this will be printed too")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}

}
