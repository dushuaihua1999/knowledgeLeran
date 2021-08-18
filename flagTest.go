package main

import (
	"flag"
	"fmt"
	"strings"
)

//自定义一个类似Int，Bool等的Flag命令行解析类型
type NamesFlag struct {
	Names []string
}

func (n *NamesFlag) GetNames() []string {
	return n.Names
}

//Set方法确保相关命令行选项没有被设置,之后，获取输入并使用strings.Spilt()函数来分隔参数。
//最后，参数被保存在NamesFlag结构的Names字段。
func (n *NamesFlag) Set(v string) error {
	if len(n.Names) > 0{
		return fmt.Errorf("Cannot use names flag more than once!")
	}
	names := strings.Split(v,",")
	for _,item := range names {
		n.Names = append(n.Names, item)
	}
	return nil
}

func (n *NamesFlag) String() string {
	return fmt.Sprint(n.Names)
}

func main() {
	//自定义Flag 操作数据类型
	var manyNames NamesFlag
	minusK := flag.Int("k:",0,"An int")
	minusO := flag.String("o","dushuaihua","The name")
	flag.Var(&manyNames,"names","Comma-separated list")
	flag.Parse()

	fmt.Println("-K",*minusK)
	fmt.Println("-O:",*minusO)



	for i,item := range manyNames.GetNames(){
		fmt.Println(i,item)
	}
	fmt.Println("Remaing command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index,": ",val)
	}
}
