package mail

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

func Sendmail(name string, address string, title string, body string) {
	// Set up authentication information.
	smtpserver := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		"coinforecasting@gmail.com",
		"aozjokifvsxhuign",
		smtpserver,
	)

	from := mail.Address{Name: "Coin Forecasting", Address: "coinforecasting@gmail.com"}
	to := mail.Address{Name: name, Address: address}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpserver+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
}
