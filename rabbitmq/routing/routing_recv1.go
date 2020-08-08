package main

import "go_devel/rabbitmq/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_one")
	imoocOne.RecieveRouting()
}
