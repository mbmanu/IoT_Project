package smtp

import (
	"net/smtp"
	"os"
)

func SendEmail(to []string, subject, body string) error {
	from := os.Getenv("EMAIL_ADDR")
	password := os.Getenv("EMAIL_PSWD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")

	address := host + ":" + port

	msg_to := func(to ...string) string {
		var rv string
		for _, t := range to {
			rv += t + " "
		}
		return rv
	}

	message := []byte("To: " + msg_to(to...) + "\r\n" + "Subject: " + subject + "\r\n" + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	return err
}
