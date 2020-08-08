package main

import "go_devel/rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.ConsumeSimple()
}
