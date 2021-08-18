package main

import "fmt"

func f(slice []int) {
	slice = append(slice, 12, 14)
}

func main()  {
	slice := make([]int,10)
	f(slice)
	fmt.Println(slice)
}
