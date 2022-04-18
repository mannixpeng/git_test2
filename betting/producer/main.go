package main

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@172.16.130.28:5672/app")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"user.betting.direct_test", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//bettingStr := `{"agent_id": 1, "application_id": 1001, "bucket": 0, "channel_id": 3,"epoch": 1649823085625,"gamelog": {"account":"xcv1","activity_bonus":0,"agent_id":1,"all_bets":40000,"all_bonus":0,"all_dpay":-40000,"channel_id":3,"client_ip":"172.16.130.112","currency_id":2,"detail":{},"device_id":1,"effect_dama":40000,"end_chips":20080,"epoch":1649823085625,"game_id":3000602,"gt_id":30,"jp_bonus":0,"jp_contri":0,"log_id":9571368719044609,"md5":"","play_id":0,"revenue":0,"room_id":0,"room_type":1,"start_chips":60080,"third_account":"36563941067841537","third_type":1,"top_agent_id":1,"type_id":1,"user_id":36563941067841537,"user_name":"xcv1"},"gtid": 30,"key": "","parent_id": 0,"platform_id": 3}`
	bettingStr := `{"id":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","log_id":18574726104317954,"platform_id":3,"application_id":1001,"channel_id":10011001,"agent_id":0,"user_id":331,"month":202204,"date":20220415,"is_active":0,"betting_time":0,"betting_out_time":0,"betting_duration":0,"betting_amount":250,"betting_count":0,"win_lose_amount":150,"tax_amount":0,"rebate_amount":150,"income_rate":0,"win_lose_tax_rate":0,"win_lose_betting_rate":0,"all_bets":250,"all_bonus":400,"jp_bonus":0,"jp_contri":0,"betting_number":0}`
	wait := sync.WaitGroup{}
	wait.Add(10000*1000000000)
	for i := 0; i < 10000; i++ {
		go func(wait *sync.WaitGroup) {
			for i := 0; i< 1000000000; i++ {
				err = ch.Publish(
					"user.betting.direct_test",         // exchange
					//severityFrom(os.Args), // routing key
					"betting.count.test", // routing key
					false,                 // mandatory
					false,                 // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(bettingStr),
					})
				failOnError(err, "Failed to publish a message")

				log.Printf(" [x] Sent %s", bettingStr)
				wait.Done()
			}
		}(&wait)
	}
	wait.Wait()
	select {

	}
}

