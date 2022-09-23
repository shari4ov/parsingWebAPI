package apis

import (
	dto "api/DTO"
	"api/storage"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func JsonToObject(jsonData []byte) dto.CreateNews {
	var n dto.CreateNews
	_ = json.Unmarshal(jsonData, &n)
	return n
}
func SaveToDB(ch chan dto.CreateNews) {
	news := <-ch
	fmt.Println(news)
	db := storage.OpenConnection()
	sqlStatement := `INSERT INTO news (title,news_id,href) VALUES($1,$2,$3);`
	_, err := db.Exec(sqlStatement, news.Title, news.News_id, news.Href)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
func GetLastNews() (dto.News, error) {
	var n dto.News
	db := storage.OpenConnection()
	sqlLastRow := `SELECT * FROM news ORDER BY id DESC LIMIT 1;`
	row := db.QueryRow(sqlLastRow)
	err := row.Scan(&n.Id, &n.Title, &n.News_id, &n.Href)
	fmt.Println(n, "12121````````")
	if err != nil {
		fmt.Println(err)
		return n, err
	}
	db.Close()
	return n, nil
}
func SendLastNews(news dto.News) error {
	conn, err := amqp091.Dial("amqp://admin:admin@localhost:5672")
	fmt.Println(err)
	amqpChannel, err := conn.Channel()
	fmt.Println(err)
	defer conn.Close()
	defer amqpChannel.Close()
	q, error := amqpChannel.QueueDeclare(
		"lastNews.queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if error != nil {
		fmt.Println(error)
	}
	amqpChannel.QueueBind(
		q.Name,
		"",
		"amq.fanout",
		false,
		nil,
	)
	news_Marshalled, _ := json.Marshal(news)
	msg := amqp091.Publishing{
		Body: []byte(news_Marshalled),
	}
	amqpChannel.Publish(
		"",
		q.Name,
		false,
		false,
		msg,
	)
	return nil
}
func CreateNews() {

	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	HandleError(err, "Line 87")
	defer conn.Close()
	ch, err := conn.Channel()
	HandleError(err, "Line 92")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"news.queue",
		true,
		false,
		false,
		false,
		nil,
	)
	HandleError(err, "Line 104")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	HandleError(err, "Line 111")
	runtime.GOMAXPROCS(3)
	forever := make(chan bool)
	channel := make(chan dto.CreateNews, 1)
	go func() {
		for d := range msgs {
			channel <- JsonToObject(d.Body)
			SaveToDB(channel)
			time.Sleep(time.Second * 3)
		}

	}()
	<-forever
}
func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(err, msg)
	}
}
