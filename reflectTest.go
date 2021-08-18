package main

import (
	"fmt"
	"reflect"
)

func main()  {
	t := 18
	refT := reflect.TypeOf(t)
	fmt.Println(refT.String())
}
