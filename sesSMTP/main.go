package main

import (
	"log"
	"net/smtp"
)

// variables to make ExamplePlainAuth compile, without adding
// unnecessary noise there.
var (
	from       = "gopher@mail.example.com"
	msg        = []byte("dummy message")
	recipients = []string{"abc@xyz.com"}
)

func main() {
	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "email-smtp.ap-northeast-1.amazonaws.com"
	auth := smtp.PlainAuth("gopher@mail.2xtree.xyz", "AKxxxxxx", "yyyyyyyyyyyyyyyyyyyyyyyy", hostname)

	err := smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}