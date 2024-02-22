package main

import "fmt"

func swap(e, f *int) {
	fmt.Println("In Swap : ", &e, " : ", &f)
	fmt.Println("In Swap : ", e, " : ", f)
	fmt.Println("In Swap : ", *e, " : ", *f)
	temp := *e
	*e = *f
	*f = temp
}

func swapNormal(i, j int) (int, int) {
	temp := i
	i = j
	j = temp
	return i, j
}

func main() {
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
	fmt.Println(&x, " : ", &y)
	swap(&x, &y)
	fmt.Println(x, y) // 10, 5

	i := 50
	j := 100
	fmt.Println(i, j) // 50, 100
	var m, n int
	m, n = swapNormal(i, j)
	fmt.Println(m, n) // 100, 50
}
