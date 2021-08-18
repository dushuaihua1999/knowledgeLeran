package main

import (
	"fmt"
	"strings"
)

func main()  {
	var b strings.Builder
	//for i := 0; i < 3; i++ {
	//	fmt.Fprintf(&b,"%d...",i)
	//}
	b.WriteString("ignition")
	fmt.Println(b.String())
}
