package service

import (
	"net/smtp"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/internal/config"
)

type emailService struct {
	cnf *config.Config
}

func NewEmail(cnf *config.Config) domain.EmailService {
	return &emailService{
		cnf: cnf,
	}
}

func (e emailService) Send(to, subject, body string) error {
	myEmail := "atok@gmai.com"
	msg := []byte("" +
		"From: <" + myEmail + ">\n" +
		"To:" + to + "\n" +
		"Subject: " + subject + "\n" +
		"\n" +
		body + "\n")

	return smtp.SendMail(e.cnf.Email.Host+":"+e.cnf.Email.Port, nil, myEmail, []string{to}, []byte(msg))
}
