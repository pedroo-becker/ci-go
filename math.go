package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(notReachable())

}

func soma(a int, b int) int {
	return a + b
}

func notReachable() int {
	if true {
		return 1
	}

	return 0
}
