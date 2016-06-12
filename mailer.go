package mailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"mime/quotedprintable"
	"net"
	net_mail "net/mail"
	"net/smtp"

	"time"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type Config interface {
	SmtpUser() string
	SmtpPassword() string
	SmtpHost() string
	SmtpPort() int
	Tls() bool
	TlsSkipVerify() bool
	Timeout() time.Duration
}

type Message interface {
	Sender() string
	Recipient() string
	Content() string
	Subject() string
}

type mailer struct {
	config Config
}

type Mailer interface {
	Send(message Message) error
}

func New(config Config) *mailer {
	m := new(mailer)
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

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = "=?UTF-8?Q?" + QuoteString(message.Subject()) + "?="
	headers["Content-Type"] = `text/plain; charset="utf-8"`
	headers["Content-Transfer-Encoding"] = `quoted-printable`

	content := ""
	for k, v := range headers {
		content += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	content += "\r\n" + QuoteString(message.Content())

	logger.Tracef("connect to %s", servername)
	conn, err := createConn(servername, s.config.Tls(), s.config.TlsSkipVerify(), s.config.Timeout())
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

func QuoteString(s string) string {
	w := bytes.NewBufferString("")
	qw := quotedprintable.NewWriter(w)
	qw.Write([]byte(s))
	qw.Close()
	return string(w.Bytes())
}

func createConn(servername string, tlsActive bool, tlsSkipVerify bool, timeout time.Duration) (net.Conn, error) {
	dailer := &net.Dialer{Timeout: timeout}
	if tlsActive {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: tlsSkipVerify,
			ServerName:         servername,
		}
		logger.Debugf("open tls connection to: %s", servername)
		return tls.DialWithDialer(dailer, "tcp", servername, tlsconfig)
	}
	logger.Debugf("open non tls connection to: %s", servername)
	return dailer.Dial("tcp", servername)
}
