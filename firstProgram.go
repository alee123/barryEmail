package main

import (
	"fmt"
	"github.com/pcrawfor/gopostal"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

type messageText struct {
	Subject  string
	TextBody string
	HtmlBody string
}

func mailer(c chan *messageText, input string) {
	mailer := gopostal.NewGmailMailer("clypdMonster", "clypdsecret")
	for msg := range c {
		toSend := mailer.NewMessage(input, "Barry", msg.Subject, msg.TextBody, msg.HtmlBody)
		//toSend.AddTo("richard@clypd.com")
		err := mailer.Send(*toSend)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}
}

func convertToMessage(subject string, body string, htmlbody string, c chan *messageText) {
	msg := &messageText{
		Subject:  subject,
		TextBody: body,
		HtmlBody: htmlbody,
	}
	c <- msg
}

func main() {
	fmt.Println("Who are you sending email to?")
	var input string
	fmt.Scanf("%s", &input)
	subjects := []string{"Hello", "is", "it", "me", "you're", "looking", "for"}

	fmt.Println("Sending...")
	c := make(chan *messageText)

	go mailer(c, input)

	for _, s := range subjects {
		convertToMessage(s, "Hello", "<img height='200' src='https://s3.amazonaws.com/uploads.hipchat.com/50956/344753/NxheV2TEJbWXqUW/Barry_2.jpg'> </img>", c)
		wg.Add(1)
		time.Sleep(time.Millisecond * time.Duration(10000))
	}

	wg.Wait()
}
