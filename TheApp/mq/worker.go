package main

import (
	"bytes"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"


	"github.com/streadway/amqp"
)

type User struct {
	//Id      int   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Password string `json:"password"`
	Role string `json:"role"`
	Operation string `json:"opertion"`
}

func (b *User) TableName() string {
	return "user"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//var err error
	db, err := gorm.Open(sqlite.Open("C:\\Users\\aandrianto\\TrainingGo\\g2-go\\TheApp\\thedb.db"), &gorm.Config{})
	if err != nil{
		panic("Cannot connect to DB")
	}

	db.AutoMigrate(User{})
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {

		for d := range msgs {
			//var user Models.User
			user := User{}
			log.Printf("Raw message: %s", d.Body)
			json.Unmarshal([]byte(d.Body), &user)
			//switch
			log.Printf("Received a message: %s", user.Name)
			//Models.CreateUser(user)
			db.Create(user)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}