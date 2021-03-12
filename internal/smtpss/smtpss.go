package smtpss

import (
	"net/smtp"
)

func PlainAuth(username string, password string, hostname string, from string, msg []byte, recipients []string) error {
	// hostname is used by PlainAuth to validate the TLS certificate.
	auth := smtp.PlainAuth("", username, password, hostname)

	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return err
}
