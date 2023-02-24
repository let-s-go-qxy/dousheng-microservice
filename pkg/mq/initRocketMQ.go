package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

var (
	producer1 rocketmq.Producer
	consumer1 rocketmq.PushConsumer
)

func InitRM() (rocketmq.Producer, rocketmq.PushConsumer, error) {
	var err error
	producer1, err = rocketmq.NewProducer(
		producer.WithVIPChannel(false),
		producer.WithSendMsgTimeout(time.Second*6),
		producer.WithNameServer([]string{"127.0.0.1:9876"}), // 接入点地址
		producer.WithRetry(100),                             // 重试次数
		producer.WithGroupName("CommentProductGroup"),       // 分组名称
	)
	if err != nil {
		fmt.Println("生产者创建失败：", err.Error())
		return nil, nil, err
	}
	consumer1, err = rocketmq.NewPushConsumer(consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("CommentConsumerGroup"),
		//consumer.WithConsumerOrder(true),
	)
	if err != nil {
		fmt.Println("消费者创建失败：", err.Error())
		return nil, nil, err
	}
	return producer1, consumer1, nil
}

func SendSyncMessage(message string) error {
	err := producer1.Start()
	// 发送的消息
	msg := &primitive.Message{
		Topic: "testTopic01",
		Body:  []byte(message),
	}
	// 发送消息
	_, err = producer1.SendSync(context.Background(), msg)
	if err != nil {
		fmt.Println("发送消息失败：", err.Error())
		return err
	}
	return nil
}

// SubscribeMessage 订阅并消费
func SubscribeMessage(fn func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error { // 订阅topic
	err := consumer1.Subscribe("testTopic01", consumer.MessageSelector{}, fn)
	if err != nil {
		fmt.Printf("subscribe message error: %s\n", err.Error())
		return err
	}
	// 启动consumer
	err = consumer1.Start()
	if err != nil {
		fmt.Printf("consumer start error: %s\n", err.Error())
		return err
	}
	return nil
}
