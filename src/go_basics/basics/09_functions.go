package go_basics

import (
	"fmt"
)

func Function() {
	fmt.Println("1. Each application has a main function as an entry point. It takes no arguments and returns no values.")
	fmt.Println("2. Each function is start with 'func' keyword, and the name of function determines if this is a global function or private function depending on the first letter.")
	fmt.Println("3. The function name is followed by a list of parameters in parentheses. The type comes after the variable name. If there are no parameters, the parentheses are empty. The return type comes after the parameters.")
	fmt.Println("4. The function body is enclosed in curly braces. The body can be a single statement or a block of statements. The first curly brace must be on the same line as the function declaration. The last curly brace must be on a line by itself.")
	fmt.Println("5. The return statement is used to return a value from a function. If the return type is not specified, the return type is 'void'. If there is no return statement, the return type is 'nil'.")
	fmt.Println("6. The function can be called from anywhere in the program. The function call is similar to the function declaration. The function name is followed by a list of arguments in parentheses. The arguments must match the parameters in number, order, and type.")
	fmt.Println("7. The function can return multiple values. The return type is a comma-separated list of types enclosed in parentheses. The return statement is a comma-separated list of expressions enclosed in parentheses. The expressions must match the return types in number, order, and type.")
	fmt.Println("8. The parameters can be passed by value or by reference. The value is copied to the parameter. The reference is a pointer to the original value. The reference is more efficient.The value is safer. The value cannot be changed by the function. The reference can be changed by the function.")

	for i := 0; i < 5; i++ {
		message("Hello Go!", i)
	}
	
	s := add(1, 2, 3, 4, 5)
	fmt.Println(s)

	fmt.Println(sum(1, 2))

	result, err := divide(5.6, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}  // this is a typical way to handle errors in Go.
	fmt.Println(result)

	// anonymous function
	func() {
		fmt.Println("Anonymous function")
	}() // () is used to call the anonymous function

	f := func() {
		fmt.Println("Anonymous function assigned to a variable")
	}
	f()

	// function as a variable
	var div func(float64, float64) (float64, error)
	div = divide
	result, err = div(5.6, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	
	// function as a type
	type calc func(int, int) int // define a type
	var addition calc = func(a int, b int) int {  // define a function
		return a + b
	}
	fmt.Println(addition(1, 2))

	// function as a method
	p := person{name: "John", age: 25}
	p.speak()

}	

func message(msg string, idx int) {
	fmt.Println(msg, idx)
}


func add(values ...int) int {  
	// ... means that the function can take any number of arguments
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	// fmt.Println("The summe is", result)
	return result
}

func sum(a, b int) (result int) {  // result is a named return value
	result = a + b
	return  // return without arguments returns the named return value
}

func divide(a, b float64) (float64, error) {  
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}  //use an error instead of panic the program.
	return a / b, nil
}

// method
type person struct {
	name string
	age  int
}

func (p person) speak() {  // p is a value receiver
	fmt.Println(p.name, "is", p.age, "years old")
}
