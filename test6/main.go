package  main

import (
   "fmt"
   "test6/cmd"
   "time"
)

func main(){
   //testmq()
   //cmd.TestKafkaConsumer()
   //cmd.TestNsqProducer()
   cmd.TestNsqConsumer()
   //cmd.TestOrder()
   time.Sleep(10*time.Second)
   fmt.Println("hello","world")
}