package test

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestTopic(t *testing.T) {
	// 1. 创建主题，这一步可以省略，在send的时候如果没有topic，也会进行创建。
	//CreateTopic("testTopic01")
	rlog.SetLogLevel("error")
	// 2.生产者向主题中发送消息
	//SendSyncMessage("------------")
	//SubscribeMessage()
	go Subs()
	// 3.消费者订阅主题并消费
	go Send()
	time.Sleep(time.Second * 10)
}

func Send() {
	for i := 0; i < 100; i++ {
		SendSyncMessage("消息" + strconv.Itoa(i) + time.Now().String())
	}
}

func Subs() {
	SubscribeMessage()
}

func CreateTopic(topicName string) {
	endPoint := []string{"127.0.0.1:9876"}
	// 创建主题
	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(endPoint)))
	if err != nil {
		fmt.Printf("connection error: %s\n", err.Error())
	}
	err = testAdmin.CreateTopic(context.Background(),
		admin.WithBrokerAddrCreate("127.0.0.1:10911"),
		admin.WithTopicCreate(topicName),
		admin.WithWriteQueueNums(1),
	)
	if err != nil {
		fmt.Printf("createTopic error: %s\n", err.Error())
	}
}

func SendSyncMessage(message string) {
	p, err := rocketmq.NewProducer(
		producer.WithVIPChannel(false),
		producer.WithSendMsgTimeout(time.Second*6),
		producer.WithNameServer([]string{"127.0.0.1:9876"}), // 接入点地址
		producer.WithRetry(2),                               // 重试次数
		producer.WithGroupName("CommentProductGroup"),       // 分组名称
	)
	if err != nil {
		fmt.Println("生产者创建失败：", err.Error())
		return
	}
	err = p.Start()
	if err != nil {
		fmt.Println("生产者开始失败：", err.Error())
		return
	}
	// 发送的消息
	msg := &primitive.Message{
		Topic: "testTopic01",
		Body:  []byte(message),
	}
	// 发送消息
	_, err = p.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Println("发送消息失败：", err.Error())
		return
	}
	fmt.Println("发送成功")
	defer func(p rocketmq.Producer) {
		err = p.Shutdown()
		if err != nil {
			fmt.Println("生产者关闭失败：", err.Error())
		}
	}(p)
}

// SubscribeMessage 订阅并消费
func SubscribeMessage() {
	// 创建一个consumer实例
	c, err := rocketmq.NewPushConsumer(consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("CommentConsumerGroup"),
		//consumer.WithConsumerOrder(true),
	)

	// 订阅topic
	err = c.Subscribe("testTopic01", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("订阅到的消息 : %v -> %v \n", string(msgs[i].Message.Body), i)
			//time.Sleep(time.Second * 3)
			//fmt.Println("处理完了")
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Printf("subscribe message error: %s\n", err.Error())
	}

	// 启动consumer
	err = c.Start()

	if err != nil {
		fmt.Printf("consumer start error: %s\n", err.Error())
		os.Exit(-1)
	}

	//defer func(c rocketmq.PushConsumer) {
	//	err = c.Shutdown()
	//	if err != nil {
	//		fmt.Printf("shutdown Consumer error: %s\n", err.Error())
	//	}
	//}(c)
}
