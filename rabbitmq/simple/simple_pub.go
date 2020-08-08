package main

import (
	"fmt"
	"rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	fmt.Println("发送成功！")
}
