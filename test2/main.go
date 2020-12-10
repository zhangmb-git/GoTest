package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){
	//TestChannel()
	wg:=sync.WaitGroup{}
	//TestTicker()
	client :=&Client{
		tickerDone: make(chan struct {},1),
		ticker:     time.NewTicker(5 * time.Second),
	}
    wg.Add(2)
	go KeepAlive(client,&wg)
	go CloseTime(client,&wg)
	//client := cli.NewClient()
	//client.Start()
	fmt.Println("test end")
	//select {}
    //time.Sleep(30)
    wg.Wait()
}