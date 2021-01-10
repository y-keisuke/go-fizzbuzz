package main

import (
	"fizzbuzz/fizzbuzz"
	"fmt"
)

func main() {
	fmt.Println("Let's FizzBuzz")

	for i := 1; i <= 30; i++ {
		fmt.Println(fizzbuzz.Run(i))
	}
}
