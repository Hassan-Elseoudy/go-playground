package main

import "fmt"

func main() {
	var greetings = `Greetings and
					 "salutations!""`
	var flag = true
	// Complex number
	x := 1 + 2i
	y := 2 + 3i
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	// constants
	const name = "Hassan Elseoudy"
	fmt.Println(name)
	// arrays
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// slices
	var slice = []int{1, 2, 3, 4, 5}
	var slice2 = []int{6, 7, 8, 9, 10}
	// append to slice
	slice = append(slice, 6)
	// append slice to slice
	slice = append(slice, slice2...)
	// print capacity of slice, and length of slice
	fmt.Println(cap(slice), len(slice))
	fmt.Println(slice)
	fmt.Println(flag)
	fmt.Println(greetings)
}
