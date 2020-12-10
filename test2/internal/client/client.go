package client

import (
	"log"
	"net"

)

type Options struct {
	headerLen            int           // 数据包头部大小,默认值是2
	ServerAddr           string
}

type  Client struct {
	options           *Options     // 服务参数
	conn              net.Conn     //连接
}


func   NewClient() (* Client) {
	 options := &Options{
		 2,
		 "10.0.72.202:8600",
	 }

	 conn,err:=net.Dial("tcp",options.ServerAddr)
	 if err != nil {
	 	log.Fatalf("connect [%s] failed,%s",options.ServerAddr ,err.Error())
		return  nil
	 }

     client :=&Client{
        options:options,
        conn:conn,
	 }
	 return  client
}

func  ( this * Client)  Start(){
    go CoConnLoop(this)
}

func (this*  Client) Stop(){
	this.conn.Close()
}

func  CoConnLoop(c*  Client){
	codec := NewCodec(c.conn)

	//循环读
	go func() {
		for {
			_, err := codec.Read()
			if err != nil {
				log.Println(err)
				return
			}
			for {
				bytes, ok, err := codec.Decode()
				// 解码出错，需要中断连接
				if err != nil {
					log.Println(err)
					return
				}
				if ok {
					log.Println(string(bytes))
					continue
				}
				break
			}
		}
	}()

}