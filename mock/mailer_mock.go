package mock

import "github.com/bborbe/mail"

type mailer struct {
	Error   error
	Message mail.Message
	Counter int
}

func New() *mailer {
	m := new(mailer)
	m.Counter = 0
	return m
}

func (m *mailer) Send(message mail.Message) error {
	m.Message = message
	m.Counter++
	return m.Error
}
