package main

import (
	"fmt"
	"math/rand"
)

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
	fmt.Println(slice)
	var slice2 = []int{6, 7, 8, 9, 10}
	// append to slice
	slice = append(slice, 6)
	// append slice to slice
	slice = append(slice, slice2...)
	// print capacity of slice, and length of slice
	fmt.Println(cap(slice), len(slice))
	// slice literal
	slice3 := []int{1, 2, 3, 4, 5}
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	// slices share the same memory
	slice4 := slice3
	slice4[0] = 100
	fmt.Println(slice3)
	// copy slices
	slice5 := make([]int, len(slice3))
	copy(slice5, slice3)
	slice5[0] = 200
	fmt.Println(slice3)
	fmt.Println(slice5)
	// maps
	var m = map[string]int{}
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(m["one"])
	v, ok := m["three"]
	fmt.Println(v, ok)
	// delete from map
	delete(m, "one")
	fmt.Println(m)

	// structs
	type person struct {
		name string
		age  int
	}
	p := person{name: "Hassan", age: 25}
	p.name = "Semsemawy"
	fmt.Println(p)

	// for loop to calculate 2^1 until 2^30 and sum them
	sum := 0
	for i := 0; i < 30; i++ {
		sum += 1 << i
		fmt.Println(1 << i)
	}
	fmt.Println(sum)

	fmt.Println(flag)
	fmt.Println(greetings)

	// example 4-1. Shadowing variables
	va := 10
	if va > 5 {
		fmt.Println(va)
		va := 5
		fmt.Println(va)
	}
	fmt.Println(va)

	// example 4-2. Shadowing with multiple variables
	vb := 10
	if vb > 5 {
		vb, vc := 5, 20
		fmt.Println(vb, vc)
	}
	fmt.Println(vb)

	// example 4-5 if and else
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("n is 0")
	} else if n > 0 {
		fmt.Println("n is greater than 0")
	} else {
		fmt.Println("n is less than 0")
	}

}
