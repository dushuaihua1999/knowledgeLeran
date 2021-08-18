package main

import (
	"fmt"
	"time"
)
//指针通道可以修改原值
//Counter代表计数器的类型
type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func main()  {
	syncChan := make(chan struct{}, 2)

	go func() {
		for {
			if elem, ok := <-mapChan; ok{
				counter := elem["count"]
				counter.count++
			}else{
				break
			}
		}
		fmt.Println("Stopped [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]*Counter{
			"count" : &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. count:%v [sender]\n", countMap["count"].count,countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
