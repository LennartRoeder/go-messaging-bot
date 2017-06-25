package http

import (
	"fmt"

	"strconv"

	"time"

	"github.com/headzoo/surf"
)

var bow = surf.NewBrowser()

const url string = "https://10minut.com.pl/"

func CreateEmail() string {
	err := bow.Open(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("CreateEmail:", bow.Title())

	email := bow.Find("input#inputEmail3")

	return email.Nodes[0].Attr[3].Val
}

func ActivateEmailAccount() {
	fetchEmail(1)
	activate()
}

func fetchEmail(id int) {
	time.Sleep(time.Second * 10)
	fmt.Println("open:", url+"?mail="+strconv.Itoa(id))

	err := bow.Open(url + "?mail=" + strconv.Itoa(id))
	if err != nil {

		// ToDo: retry until email arrived

		panic(err)
	}

	fmt.Println("fetchEmail:", bow.Title())

	activate := bow.Find("td.button > a")

	fmt.Println(activate.Nodes[0])
	fmt.Println(activate.Nodes[0].Attr[3].Val)
}

func activate() {
	// ToDo: implement
}
