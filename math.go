package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(notReachable())

}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func foo(x, y int) {
	switch x {
	case 0:
		switch y { // Noncompliant; nested switch
		case 1:
			fmt.Println(notReachable())
		}
	case 1:
		fmt.Println(notReachable())
	default:
		fmt.Println(notReachable())
	}
}

func print(a int, b int) int {
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	fmt.Println(soma(a, b))
	return 0
}

func notReachable() int {
	if true {
		return 1
	}

	return 0
}
