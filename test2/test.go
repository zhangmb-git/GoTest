package main

import (
	"fmt"
	"sync"
	"time"
)


func TestChannel() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3

	// 关闭函数非常重要,若不执行close(),那么range将无法结束,造成死循环
	// close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func TestTicker(){


}

type Client struct {
	tickerDone  chan struct{}
	ticker       *time.Ticker
}

func  KeepAlive(it *Client,wg * sync.WaitGroup){
	fmt.Println("KeepAlive...")
	defer wg.Done()
	for {
		select {
		case <-it.tickerDone:
			fmt.Println("exit keepAliveRoutine")
			fmt.Println("exit keepAliveRoutine")
			return
		case <-it.ticker.C:
			fmt.Println("ticker...")
		}
	}
	fmt.Println("KeepAlive Over")
}

func CloseTime(it *Client,wg * sync.WaitGroup){
	defer  wg.Done()
	time.Sleep(11*time.Second)
	//if it.tickerDone != nil {
	//	close(it.tickerDone)
	//	it.tickerDone = nil
	//}
	if it.ticker != nil {
		it.ticker.Stop()
		fmt.Println("close ticker...")
		//close(it.tickerDone)
		//it.ticker = nil
		//time.Sleep(6*time.Second)
		//close(it.tickerDone)


	}
	fmt.Println("closetime over....")

}

