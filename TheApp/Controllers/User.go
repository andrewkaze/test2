package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"theapp/Models"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(user *Models.User,operation string) string {

	s:=`{ "name" : `+ `"`+user.Name+`"` + `,"email" : ` + `"` + user.Email +`"` + `,"address" : ` +`"`+user.Address+`"`+ `,"operation" : `+`"` +operation+`"`+` }`

	return s
}

func SendMq (user *Models.User,operation string) error{
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

	body := bodyFrom(user,operation)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	return err
}

func GetUsers(c *gin.Context)  {
	var user[]Models.User
	err := Models.GetAllUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		format := c.DefaultQuery("format", "json")
		if format == "json"{
			c.JSON(http.StatusOK, user)
		}else{
			c.XML(http.StatusOK, user)
		}
	}
}

func CreateUser(c *gin.Context){
	var user Models.User
	c.BindJSON(&user)
	err := SendMq(&user,"createUser")
	//err := Models.CreateUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil{
		c.JSON(http.StatusNotFound, user)
	}

	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}