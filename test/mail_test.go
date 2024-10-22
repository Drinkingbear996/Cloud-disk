package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <Drinkingbear996@gmail.com>"
	e.To = []string{"huan2846@umn.com"}
	e.Subject = "Verification code Test"
	e.HTML = []byte("Your code is<h1>123456</h1>")

	// Ensure you are using the correct app-specific password for Gmail
	err := e.SendWithTLS(
		"smtp.gmail.com:465",
		smtp.PlainAuth("", "Drinkingbear996@gmail.com", define.Gmailpwd, "smtp.gmail.com"),
		&tls.Config{ServerName: "smtp.gmail.com"},
	)

	if err != nil {
		t.Fatal(err)
	}
}
