package http

import (
	"fmt"

	"math/rand"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/private/go-messaging-bot/template"
)

func RegisterAccount(url string, email string) template.Account {
	register(url, email)
	ActivateEmailAccount()

	return template.Account{"", "", ""}
}

func register(url string, email string) {
	fmt.Println("register account: ", url, email)

	err := bow.Open(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("register:", bow.Title())

	fm, err := bow.Form("form#application_register")
	if err != nil {
		panic(err)
	}

	// gender
	err = fm.Input("gender", "2")
	if err != nil {
		panic(err)
	}

	// email
	err = fm.Input("email", email)
	if err != nil {
		panic(err)
	}

	// Day of birth (1-27)
	dayOfBirth := strconv.Itoa(getRandom(27) + 1)
	err = fm.Input("dob_day", dayOfBirth)
	if err != nil {
		panic(err)
	}

	// Month of birth (1-12)
	monthOfBirth := strconv.Itoa(getRandom(12) + 1)
	err = fm.Input("dob_month", monthOfBirth)
	if err != nil {
		panic(err)
	}

	// year of birth (87-97)
	yearOfBirth := strconv.Itoa(getRandom(10) + 1988)
	err = fm.Input("dob_year", yearOfBirth)
	if err != nil {
		panic(err)
	}

	// submit
	err = fm.Submit()
	if err != nil {
		panic(err)
	}

	// show registration errors
	bow.Dom().Find(".dtyrd-form-error-box").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func getRandom(max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	return rnd.Intn(max)
}
