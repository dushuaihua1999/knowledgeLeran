package main

import "github.com/PuerkitoBio/goquery"

type FilePipeLine struct {
	dir string
}

func NewFilePipeLine(dir string) *FilePipeLine {
	return &FilePipeLine{dir: dir}
}

func (p *FilePipeLine) Process(doc *goquery.Document)  {
	//文件写入实现
}
