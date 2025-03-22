package main

import "fmt"

func append1(in []int64, value int64) {
	in = append(in, value)
}

func mutate1(in []int64, idx, value int64) {
	in[idx] = value
}

func main() {
	slice := make([]int64, 0, 4)
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice) //1, 2
	append1(slice, 3)
	slice = slice[:3] //
	fmt.Println(slice)
	mutate1(slice, 3, 4)
	fmt.Println(slice)
}
