package go_basics

import (
	"fmt"
)

func Pointer() {
	// poiner arithmetic is not allowed in Go
	// * is the pointer operator
	// & is the address-of operator
	fmt.Println("Pointer")
	a := 42
	b := a  // b is a copy of a
	c := &a  // c is a pointer to a
	fmt.Println(a, b, c)
	a = 27  // change a
	fmt.Println(a, b, c)  // b is unchanged, c is changed

	var d int = 42
	var e *int = &d  // e is a pointer to d
	fmt.Println(d, e)
	*e = 27  // change d through e
	fmt.Println(d, e)  // d is changed, e is unchanged

	f := [3]int{1, 2, 3}  // f is an array
	g := &f[0]  // g is a pointer to f[0]
	h := &f[1]  // h is a pointer to f[1]
	fmt.Println(f, g, h)
	f[1] = 42  // change f[1]
	fmt.Println(f, g, h)  // g and h are unchanged

	i := []int{1, 2, 3}  // i is a slice
	j := i
	fmt.Println(i, j)
	i[1] = 42  // change i[1]
	fmt.Println(i, j)  // j is changed
}