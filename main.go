package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	z := new(int)
	fmt.Println(&z) // 0xZ01
	fmt.Println(z)  // 0xZ02
	fmt.Println(*z) // 0
	fmt.Println("=====")

	a := new(int)
	*a = 5
	fmt.Println(&a) // 0xA01
	fmt.Println(a)  // 0xA02
	fmt.Println(*a) // 5
	fmt.Println("=====")

	b := 10
	fmt.Println(&b) // 0xB01
	fmt.Println(b)  // 10
	//fmt.Println(*b) // error
	fmt.Println("=====")

	*a = b
	fmt.Println(&a) // 0xA01
	fmt.Println(a)  // 0xA02
	fmt.Println(*a) // 10
	fmt.Println("=====")

	b = 500
	a = &b
	fmt.Println(&b) // 0xB01
	fmt.Println(b)  // 500
	fmt.Println(&a) // 0xA01
	fmt.Println(a)  // 0xB01
	fmt.Println(*a) // 500
	fmt.Println("=====")

	x := 5
	y := 10
	fmt.Println(x, y) // 5, 10
	swap(&x, &y)
	fmt.Println(x, y) // 10, 5
}
