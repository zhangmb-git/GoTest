package cmd
import (
	"fmt"
	gonsq "github.com/nsqio/go-nsq"
	"time"

)

var nsqproducer *gonsq.Producer

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *gonsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

// NsqConsumer 消费消息
func NsqConsumer(topic, channel string,  concurrency int) {
	config := gonsq.NewConfig()

	config.LookupdPollInterval = 1 * time.Second

	consumer, err := gonsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	consumer.AddHandler(&NSQHandler{})
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		panic(err)
	}
}


func initNsqProducer() {
	var err error
	cfg := gonsq.NewConfig()
	nsqproducer, err = gonsq.NewProducer("127.0.0.1:4150", cfg)
	if nil != err {
		panic("nsq new panic")
	}

	err = nsqproducer.Ping()
	if nil != err {
		panic("nsq ping panic")
	}
}

func TestNsqProducer(){
	initNsqProducer()
	for i:=0;i<5;i++ {
		jsonstr:= []byte(fmt.Sprintf("testkey:%d",i))
		//jsonstr:=[]byte("{\"testkey\":1}")
		err := nsqproducer.Publish("testtopic",jsonstr)
		if err != nil {
			panic(err)
		}
	}

}

func TestNsqConsumer(){
	NsqConsumer("testtopic","test",2)

}


