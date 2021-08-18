package main

import (
	"container/list"
	"crypto/md5"
	"sync"
)

type QueueScheduler struct {
	queue *list.List
	locker *sync.Mutex
	listkey map[[md5.Size]byte]*list.Element
}

func NewQueueScheduler() *QueueScheduler {
	queue :=  list.New()
	locker := new(sync.Mutex)
	listkey := make(map[[md5.Size]byte]*list.Element)

	return &QueueScheduler{
		queue: queue,
		locker: locker,
		listkey: listkey,
	}
}

func (s *QueueScheduler) Pop() (string,bool) {
	s.locker.Lock()
	if s.queue.Len() <= 0 {
		return "",false
	}
	//取实体
	e := s.queue.Front()
	url,ok := e.Value.(string)
	if !ok {
		return "",false
	}
	m := md5.Sum([]byte(url))
	delete(s.listkey,m)
	s.queue.Remove(e)
	s.locker.Unlock()
	return url, true
}

//将链接放入队列中
func (s *QueueScheduler) Push(url string){
	s.locker.Lock()
	key := md5.Sum([]byte(url))
	//链接已存在
	if _,ok := s.listkey[key]; ok {
		s.locker.Unlock()
		return
	}
	e := s.queue.PushBack(url)
	s.listkey[key] = e
	s.locker.Unlock()
}
