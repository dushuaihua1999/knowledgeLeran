package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup //waitGroup阻塞的是父协程
var times int
var m sync.Mutex

func main1() {
	for {
		time.Sleep(2 * time.Second) //1.也就是一让新建的5个go程在2秒内执行完，这种方法在无法估计准确时间的情况下是不可取的。
		if times >= 5 {             //2.另一种是生产者消费者模式,利用channel
			break //3.
		}
		fmt.Println("=============第", times, "次循环=============")
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(i int) {

				log.Println("协程", i, "执行完成")
				wg.Done()
			}(i)
		}
		times++
	}
	wg.Wait()
}

func main() {
	ch := make(chan func(), 5)
	block := make(chan struct{})
	//生产者
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second * 4)
		for {
			if times >= 5 {
				break
			}
			select {
			case <- ticker.C:
				for i := 0; i < 5; i++ {
					j := i
					time := times
					f := func() {
						log.Println("第", time, "次 ", "协程", j, "执行完成")
					}
					ch <- f
				}
				times++
			case <-block:
			}
		}
		wg.Done()
		close(ch)
		log.Println("生产者结束")
	}()

	//消费者
	wg.Add(1)
	go func() {
		ticker2 := time.NewTicker(time.Second * 6)
		for {
			select {
			case f := <-ch:
				log.Println("消费者当前时间为：",time.Now())
				for f = range ch {
					f()
				}
			case <- ticker2.C:
				log.Println("消费者当前时间为：",time.Now())
				wg.Done()
				log.Println("消费者结束")
				return
			}
		}
	}()
	wg.Wait()
}
