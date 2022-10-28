package go_basics

import (
	"fmt"
	"reflect"
)

// 8. 定义一个结构体，结构体的字段必须是大写开头的，否则无法被外部包访问。
type Vertex struct {
	Name string
	Lat float64
	Long float64
}

type Animal struct {
	Name string `required max:"100"`  // 18. 声明一个结构体，结构体的字段可以有tag。
	Age int
}

type Dog struct {
	Animal  // 15. 嵌套结构体，继承Animal的字段。
	Owner string
}

func Maps_and_structs() {
	// 1. 声明一个map，key是string类型，value是int类型。
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
		"Pennsylvania": 12802503,
		"Illinois":   12801539,
		"Ohio":       11614373,
	}
	fmt.Println(statePopulations)

	a := make(map[string]int)  // 2. 使用make函数创建一个map。
	a["key"] = 10  // 3. 使用key为map赋值。
	fmt.Println(a)
	fmt.Println(a["key"])  // 4. 使用key获取map的值。

	delete(a, "key")  // 5. 使用delete函数删除map中的元素。
	fmt.Println(a)
	fmt.Println(a["key"])  // 6. 获取map中不存在的key的值，返回值类型的零值。

	b, ok := a["key"]  // 7. 使用两个变量接收map的值，第二个变量是bool类型，表示key是否存在。
	fmt.Println(b, ok)

	vertex := Vertex{"ver", 1.1, 2.2}  // 9. 声明一个结构体。
	fmt.Println(vertex)
	fmt.Println(vertex.Name)  // 10. 访问结构体的字段。

	// 11. 声明一个结构体指针。
	var vertexPointer *Vertex = &vertex
	fmt.Println(vertexPointer)
	fmt.Println(vertexPointer.Name)  // 12. 访问结构体指针的字段。

	anonymousStruct := struct {
		Name string
		Lat float64
		Long float64
	}{"anonymous", 1.1, 2.2}  // 13. 声明一个匿名结构体。
	fmt.Println(anonymousStruct)

	vertex1 := vertex  // 14. 结构体是值类型，赋值时会复制一份。
	vertex1.Name = "vertex1"
	fmt.Println(vertex)
	fmt.Println(vertex1)

	dog := Dog{Animal{"dog", 1}, "owner"}  // 16. 声明一个嵌套结构体。
	fmt.Println(dog)
	fmt.Println(dog.Name)  // 17. 访问嵌套结构体的字段。

	t := reflect.TypeOf(Animal{})  // 19. 获取结构体的类型。
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)  // 20. 获取结构体的tag。
}