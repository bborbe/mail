package mail

import (
	"crypto/tls"
	"net/smtp"

	net_mail "net/mail"

	"fmt"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type Config interface {
	SmtpUser() string
	SmtpPassword() string
	SmtpHost() string
	SmtpPort() int
}
type Message interface {
	Sender() string
	Recipient() string
	Content() string
}

type mailer struct {
	config Config
}

type Mailer interface {
	Send(message Message) error
}

func New(config Config) *mailer {
	m :=new(mailer)
	m.config = config
	return m
}

func (s *mailer) Send(message Message) error {
	logger.Debugf("sendMail to %s", message.Recipient())
	auth := smtp.PlainAuth(
		"",
		s.config.SmtpUser(),
		s.config.SmtpPassword(),
		s.config.SmtpHost(),
	)
	servername := fmt.Sprintf("%s:%d", s.config.SmtpHost(), s.config.SmtpPort())
	logger.Debugf("connect to smtp-server to %s", servername)

	from := net_mail.Address{"", message.Sender()}
	to := net_mail.Address{"", message.Recipient()}
	subj := "Monitoring Result"

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	content := ""
	for k, v := range headers {
		content += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	content += "\r\n"+message.Content()

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         servername,
	}

	logger.Tracef("connect to %s", servername)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	smtpClient, err := smtp.NewClient(conn, s.config.SmtpHost())
	if err != nil {
		return nil
	}

	err = smtpClient.Hello("localhost")
	if err != nil {
		return err
	}

	err = smtpClient.Auth(auth)
	if err != nil {
		return err
	}

	err = smtpClient.Mail(message.Sender())
	if err != nil {
		return err
	}

	err = smtpClient.Rcpt(message.Recipient())
	if err != nil {
		return err
	}

	data, err := smtpClient.Data()
	if err != nil {
		return err
	}

	logger.Tracef("write message %s", content)
	data.Write([]byte(content))

	err = data.Close()
	if err != nil {
		return err
	}

	return smtpClient.Quit()
}
