package main

import (
	"log"
	"strconv"
	"time"
)

//func main()  {
//	//1.需要几条Go程
//	//2.每条go程干多少事情负载
//	//3.什么时候阻塞，什么时候开启。
//
//}
type Spider struct {
	threadnum   uint8
	scheduler   *QueueScheduler
	downloader  *HttpDownloader
	pageprocess *PageProcess
	pipeline    *FilePipeLine
}

//创建一个爬虫引擎
func NewSpider(threadnum int, path string) *Spider {
	return &Spider{
		scheduler:   NewQueueScheduler(),
		downloader:  NewHttpDownloader(),
		pageprocess: NewPageProcess(),
		pipeline:    NewFilePipeLine(path),
		threadnum:   uint8(threadnum),
	}
}

//Run 引擎运行
func (s *Spider) Run() {
	//Process并发量
	rm := NewResourceManagerChan(s.threadnum)
	log.Println("[Spider] 爬虫运行 - 处理线程数：" + strconv.Itoa(rm.Cap()))
	for{
		url,ok := s.scheduler.Pop()
		//爬取队列为空 并且 没有Process线程 认为爬虫结束
		if ok == false && rm.Has() == 0 {
			log.Println("[Spider]爬虫运行结束")
			break
		} else if ok == false {//Process线程正在处理，可能还会有新的请求加入调度
			log.Println("[Spider] 爬取队列为空 - 等待处理")
			time.Sleep(500 * time.Millisecond)
			continue
		}
		//控制Process线程并发数量
		rm.GetOne()
		go func(url string) {
			defer rm.FreeOne()

		}(url)
	}
}

//添加请求
func (s *Spider) AddUrl(url string) *Spider {
	s.scheduler.Push(url)
	return s
}

func (s *Spider) AddUrls(urls []string) *Spider {
	for _,url := range urls {
		s.scheduler.Push(url)
	}
	return s
}

//处理请求的链接
func (s *Spider) Process(url string)  {
	//下载链接
	resp := s.downloader.Download(url)
	if resp == nil {
		/*下载失败重新加入队列中*/
		if !s.downloader.Visited(url) {
			s.scheduler.Push(url)
		}
		return
	}
	//页面处理-使用

}