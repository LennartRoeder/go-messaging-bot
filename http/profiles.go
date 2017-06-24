package http

import (
	"fmt"

	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/private/go-messaging-bot/template"
)

func GetProfiles(n int) {
	bow.Dom().Find("div#dtyrd-box-searchresults .dtyrd-user-box-content").Each(func(_ int, s *goquery.Selection) {
		age, err := strconv.Atoi(s.Find(".dtyrd-user-info-box span").First().Text())
		if err != nil {
			panic(err)
		}

		user := template.User{
			Name: s.Find(".dtyrd-user-username a").Text(),
			Age:  age,
			City: s.Find(".dtyrd-user-info-box span").First().Siblings().Text(),
			Url:  s.Find(".dtyrd-user-username a").Get(0).Attr[0].Val,
		}

		fmt.Println(user)
	})
}
