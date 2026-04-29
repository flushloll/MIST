package main

import "fmt"

func main() {
	fmt.Println("here we will perform a lot of maths")
	c1 := complex(3.4, 3.4)
	c2 := 0.54321 + 12345.6i
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)
}
