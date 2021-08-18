package main

import "github.com/PuerkitoBio/goquery"

type PageProcess struct {}

func NewPageProcess() *PageProcess {
	return &PageProcess{}
}

//返回链接函数
func (p *PageProcess) Process(d *goquery.Document) []string {
	var links []string
	//获取链接的处理代码
	return links
}