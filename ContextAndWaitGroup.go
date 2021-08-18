package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func worker(ctx context.Context)  {
	Loop:
	for{
		fmt.Println("worker is doing")
		time.Sleep(time.Second*2)
		select {
		case <- ctx.Done():
			log.Println("收到通知")
			break Loop
		default:
		}
	}
	log.Println("工人已完成工作")
	wg1.Done()
}

func main()  {
	ctx,cancel := context.WithTimeout(context.Background(), time.Second*3)        //给context上下文三秒的时间来处理事情，三秒后自动结束上下文所有线程
	wg1.Add(1)
	go worker(ctx)
	defer cancel()
	wg1.Wait()
}
