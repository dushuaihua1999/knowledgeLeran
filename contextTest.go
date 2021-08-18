package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

/*
	context的主要用途是定义Context类型和支持取消
*/

var (
	myURL string
	delay int = 5
	w     sync.WaitGroup
)

type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer w.Done()
	data := make(chan myData, 1)

	//封装一个http.Client
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", myURL, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was cancelled!")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:",err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Println("Server Response: %s\n",realHTTPData)
	}
	return nil
}
