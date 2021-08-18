package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(),"time","3s")
	ctx2 := context.WithValue(ctx,"state","谁也没整完")
	ctx3,cancel := context.WithTimeout(ctx2,3 * time.Second)
	defer cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)

	f1 := func(ctx context.Context) {
		log.Println("我是路人甲,只生产数字到ch1")
		for i := 0; i < 5; i++ {
			//通知实现加放值
			select {
			case <-ctx.Done():
				log.Println("还没做完吗？时间已经到了哦...路人甲禁止写值到ch1")
				close(ch1)
				return
			case ch1 <- i:
				time.Sleep(time.Second)
			}
		}
		close(ch1)
		log.Println("路人甲已经搞完")
		return
	}


	f2 := func(ctx context.Context) {
		log.Println("我是路人乙,要取出路人甲放进ch1的值，平方后放入ch2")
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				log.Println("时间到了，路人乙危险")
				close(ch2)
				return
			case j, ok := <-ch1:
				time.Sleep(time.Second)
				if !ok {
					log.Println("ch1管道已关闭")
				}
				ch2 <- j * j
			}
		}
		close(ch2)
		log.Println("路人乙已经搞完...")
		return
	}



	go f1(ctx3)
	go f2(ctx3)

	for i := range ch2 {
		fmt.Println("读出ch2的值：", i)
	}
}
