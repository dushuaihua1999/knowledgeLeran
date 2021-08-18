package main

import (
	"fmt"
	"knowledgeLeran/rsc.io/pdf"
)

//go语言实现目录功能
func main()  {
	genIndex("C:\\GoProject\\src\\knowledgeLeran\\readpdf\\go.pdf")
}

func genIndex(path string) string {
	is := false                //是否到目录页
	file, err := pdf.Open(path)
	if err != nil {
		panic(err)
	}
	ind := 1
	for !is {
		fmt.Println(file.Page(ind).Content())
		ind++
		if ind >= 5 {
			is = true
		}
	}
	return ""
}