package http

import (
	"fmt"
	"strings"
)

func Login(url string, email string, password string) {
	err := bow.Open(url + "login")
	if err != nil {
		panic(err)
	}

	fmt.Println(bow.Title())

	fm, err := bow.Form("form#application_login")
	if err != nil {
		panic(err)
	}

	err = fm.Input("email", email)
	if err != nil {
		panic(err)
	}

	err = fm.Input("password", password)
	if err != nil {
		panic(err)
	}

	err = fm.Submit()
	if err != nil {
		panic(err)
	}

	title := bow.Title()
	if strings.HasPrefix(title, "Umgebung - ") {
		site := strings.TrimPrefix(title, "Umgebung - ")
		fmt.Println("Login successfull:", site)
	} else {
		fmt.Println("login failed:", title)
	}
}
