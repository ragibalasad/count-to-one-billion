package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
