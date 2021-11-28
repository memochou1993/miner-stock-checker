package main

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
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
		})
		c.Visit(target)
	}
}
