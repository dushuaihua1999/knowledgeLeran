package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main()  {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { //接受操作
		//1.当前goroutine阻塞等待syncChan1传值,开始接受操作
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second...")
		//2.暂停一秒是让发送线程因为syncChan容量已满而等待接受取值而阻塞，
		//因为这时候注意点就转移到接收者这边了，不然的话，有可能发送者会一直让for循环循环完，往后执行到close(strChan)
		time.Sleep(time.Second)
		for{
			if elem, ok := <-strChan; ok{
				fmt.Println("Received: ",elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() { //发送操作
		for _,elem := range []string{"a","b","c","d"}{
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				//到此时,strChan已经填满了，可以给接收者发信号了
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 second... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan1
}
