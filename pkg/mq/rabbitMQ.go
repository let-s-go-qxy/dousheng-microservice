package mq

import (
	g "dousheng/pkg/global"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/streadway/amqp"
	"strconv"
)

type RespMessage struct {
	Id         int    `json:"id"`
	ToId       int    `json:"to_id"`
	FromId     int    `json:"from_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

// 指定MQ的Queue的Id 和strJsonMessageList 就能将和strJsonMessageList生产到对应的Queue中去
func PublishMessageCurrentToMQ(strJsonMessageList string, rabbitMQQueueId int) error {
	strRabbitMQQueueId := strconv.Itoa(rabbitMQQueueId)
	conn, err := amqp.Dial("amqp://" +
		g.RabbitMQUsername + ":" +
		g.RabbitMQPassword + "@" +
		g.RabbitMQServerAddress +
		":5672/")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	argumentsMap := map[string]interface{}{}
	argumentsMap["x-max-length"] = 1
	argumentsMap["x-overflow"] = "drop-head"
	q, err := ch.QueueDeclare(
		"message_current"+strRabbitMQQueueId, // name
		true,                                 // durable
		false,                                // delete when unused
		false,                                // exclusive
		false,                                // no-wait
		argumentsMap,                         // arguments
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strJsonMessageList),
		})
	ch.Close()
	conn.Close()
	return err
}

func PublishMessageListToMQ(strJsonMessageList string, rabbitMQQueueId int) error {
	strRabbitMQQueueId := strconv.Itoa(rabbitMQQueueId)
	conn, err := amqp.Dial("amqp://" +
		g.RabbitMQUsername + ":" +
		g.RabbitMQPassword + "@" +
		g.RabbitMQServerAddress +
		":5672/")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	argumentsMap := map[string]interface{}{}
	argumentsMap["x-max-length"] = 1
	argumentsMap["x-overflow"] = "drop-head"
	q, err := ch.QueueDeclare(
		"message_list"+strRabbitMQQueueId, // name
		true,                              // durable
		false,                             // delete when unused
		false,                             // exclusive
		false,                             // no-wait
		argumentsMap,                      // arguments
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strJsonMessageList),
		})
	ch.Close()
	conn.Close()
	return err
}

func GetRabbitMQMessageList(userId int) (respMessageList []RespMessage, err error) {
	conn, _ := amqp.Dial("amqp://" +
		g.RabbitMQUsername + ":" +
		g.RabbitMQPassword + "@" +
		g.RabbitMQServerAddress +
		":5672/")
	strUserId := strconv.Itoa(userId)
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	argumentsMap := map[string]interface{}{}
	argumentsMap["x-max-length"] = 1
	argumentsMap["x-overflow"] = "drop-head"
	q, _ := ch.QueueDeclare(
		"message_list"+strUserId, // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		argumentsMap,             // arguments
	)
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	allMessageList := []RespMessage{}
	messageList := []RespMessage{}
	//message := RespMessage{}

	go func() {
		for d := range msgs {
			//json.Unmarshal(d.Body, &message)
			json.Unmarshal(d.Body, &messageList)
			//messageList = append(messageList, message)
			allMessageList = append(allMessageList, messageList...)
		}
	}()
	err = ch.Close()
	if err != nil {
		//g.Logger.Infof("ch.Close()时发生了错误！")
	}
	err = conn.Close()
	if err != nil {
		//g.Logger.Infof("conn.Close()时发生了错误！")
	}
	return allMessageList, err
}

func GetRabbitMQMessageCurrent(userId int) (respMessageList []RespMessage, err error) {
	strUserId := strconv.Itoa(userId)
	conn, _ := amqp.Dial("amqp://" +
		g.RabbitMQUsername + ":" +
		g.RabbitMQPassword + "@" +
		g.RabbitMQServerAddress +
		":5672/")

	ch, _ := conn.Channel()

	argumentsMap := map[string]interface{}{}
	argumentsMap["x-max-length"] = 1
	argumentsMap["x-overflow"] = "drop-head"
	q, _ := ch.QueueDeclare(
		"message_current"+strUserId, // name
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		argumentsMap,                // arguments
	)
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	messageList := []RespMessage{}
	message := RespMessage{}

	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &message)
			messageList = append(messageList, message)
		}
	}()
	err = ch.Close()
	if err != nil {
		//g.Logger.Infof("ch.Close()时发生了错误！")
	}
	err = conn.Close()
	if err != nil {
		//g.Logger.Infof("conn.Close()时发生了错误！")
	}
	return messageList, err
}
