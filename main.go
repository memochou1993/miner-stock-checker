package main

import (
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"

	_ "github.com/joho/godotenv/autoload"
)

const (
	target = "https://helium.com.tw/collections/frontpage/products/2nd-copy-of-rak-hotspot-miner-donot-delete"
)

func main() {
	for range time.Tick(5 * time.Second) {
		c := colly.NewCollector()
		c.OnHTML("#product-inventory span", func(e *colly.HTMLElement) {
			availability := strings.TrimSpace(e.Text)
			log.Printf("Availability: %s", availability)
			if strings.ToUpper(availability) == strings.ToUpper("Out of stock") {
				return
			}
			notify()
		})
		c.Visit(target)
	}
}

func notify() {
	addr := "smtp.gmail.com:587"
	host := "smtp.gmail.com"
	identity := ""
	from := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("SMTP_USERNAME")
	subject := "RAK Hotspot Miner in stock now!"
	body := ""
	msg := "From:" + from + "\r\n" + "To:" + to + "\r\n" + "Subject:" + subject + "\r\n" + body
	if err := smtp.SendMail(
		addr,
		smtp.PlainAuth(identity, from, password, host),
		from,
		[]string{to},
		[]byte(msg),
	); err != nil {
		log.Println(err)
	}
}
