package main

import (
	"crypto/md5"
	"log"
	"net/http"
	"sync"
)

/*
	输入:url
	处理：
		记录下载状态
	输出：resp
*/
type HttpDownloader struct {
	locker     *sync.Mutex
	downloader map[[md5.Size]byte]bool
}

//外部新建接口
func NewHttpDownloader() *HttpDownloader {
	locker := new(sync.Mutex)
	downloaded := make(map[[md5.Size]byte]bool)
	return &HttpDownloader{
		locker:     locker,
		downloader: downloaded,
	}
}

//下载链接
//1.异常返回时记得解锁 2.resp的StatusCode码也要检验 3.没下载成功也要设置downloaded为false
func (h *HttpDownloader) Download(url string) *http.Response {
	md5Url := md5.Sum([]byte(url))
	h.locker.Lock()
	//1.判断是否已经下载过对应的连接
	if v, ok := h.downloader[md5Url]; v && ok {
		log.Println("已经下载过对应的URL")
		h.locker.Unlock()
		return nil
	}
	//2.下载对应的链接
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("获取返回值有错误", err)
		h.downloader[md5Url] = false
		h.locker.Unlock()
		return nil
	}
	//3.设置对应的状态为已经下载
	h.downloader[md5Url] = true
	h.locker.Unlock()
	log.Println("已经下载过链接了...")
	return resp
}

//判断链接是否被访问过
func (h *HttpDownloader) Visited(url string) bool {
	md5Url := md5.Sum([]byte(url))			//md5编码,因为HttpDownloader村的map的key是MD5编码过的
	h.locker.Lock()							//加锁
	sign := false                            //是否被访问过
	if v,ok := h.downloader[md5Url]; v && ok {
		log.Println("已经访问过本链接")
		sign = true
	}else{
		sign = false
	}
	h.locker.Unlock()
	return sign
}