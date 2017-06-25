package main

import "github.com/private/go-messaging-bot/http"

const dateOneUrl = "http://dateone.de/"
const dateOneEmail = "datesite@thinking-aloud.eu"
const dateOnePassword = "XxmSIvqfQV"

const datePiratUrl = "http://datepirat.de/"
const datePiratEmail = "datepirat.de@thinking-aloud.eu"
const datePiratPassword = "waUueiLAGr"

func main() {
	//os.Setenv("HTTP_PROXY", "87.79.253.219")

	email := http.CreateEmail()
	http.RegisterAccount(dateOneUrl, email)

	//http.Login(dateOneUrl, dateOneEmail, dateOnePassword)
	//http.Login(datePiratUrl, datePiratEmail, datePiratPassword)

	//http.GetProfiles(8)
}
