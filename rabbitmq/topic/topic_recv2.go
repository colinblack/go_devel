package main

import "go_devel/rabbitmq/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.*.two")
	imoocOne.RecieveTopic()
}
