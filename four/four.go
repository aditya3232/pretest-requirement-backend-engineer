package main

import "fmt"

func factorialRecursive(n uint64) uint64 {
	if n == 1 {
		return 1
	} else {
		return n * factorialRecursive(n-1)
	}
}

func main() {
	fmt.Println(factorialRecursive(10))
}
