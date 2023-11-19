package utils

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

const url string = "https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%B3%D0%BE%D1%80%D0%BE%D0%B4%D0%BE%D0%B2_%D0%A0%D0%BE%D1%81%D1%81%D0%B8%D0%B8"

func FillRedis() {
	c := colly.NewCollector()

	cities := []string{}
	//"tr > td[style=\"text-align:center\"]"
	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		cities = append(cities, text)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal("failed to visit cities resource: " + err.Error())
	}

	for _, v := range cities {
		fmt.Println(v)
	}
	fmt.Println(len(cities))
}
