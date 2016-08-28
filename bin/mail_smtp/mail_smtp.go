package main

import (
	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/bborbe/mailer"
	"github.com/bborbe/mailer/config"
	"github.com/bborbe/mailer/message"
	"github.com/golang/glog"
)

const (
	DEFAULT_HOST              = "localhost"
	DEFAULT_PORT              = 1025
	DEFAULT_TLS               = false
	DEFAULT_TLS_SKIP_VERIFY   = false
	DEFAULT_FROM              = "test@example.com"
	DEFAULT_TO                = "test@example.com"
	DEFAULT_BODY              = "Hello World\r\n"
	DEFAULT_SUBJECT           = "Test Mail"
	PARAMETER_SMTP_HOST       = "smtp-host"
	PARAMETER_SMTP_PORT       = "smtp-port"
	PARAMETER_TLS             = "smtp-tls"
	PARAMETER_TLS_SKIP_VERIFY = "smtp-tls-skip-verify"
	PARAMETER_FROM            = "from"
	PARAMETER_TO              = "to"
	PARAMETER_SUBJECT         = "subject"
	PARAMETER_BODY            = "body"
)

var (
	smtpHostPtr          = flag.String(PARAMETER_SMTP_HOST, DEFAULT_HOST, "smtp host")
	smtpPortPtr          = flag.Int(PARAMETER_SMTP_PORT, DEFAULT_PORT, "smtp port")
	smtpTlsPtr           = flag.Bool(PARAMETER_TLS, DEFAULT_TLS, "smtp tls")
	smtpTlsSkipVerifyPtr = flag.Bool(PARAMETER_TLS_SKIP_VERIFY, DEFAULT_TLS_SKIP_VERIFY, "smtp tls skip verify")
	fromPtr              = flag.String(PARAMETER_FROM, DEFAULT_FROM, "from")
	toPtr                = flag.String(PARAMETER_TO, DEFAULT_TO, "to")
	subjectPtr           = flag.String(PARAMETER_SUBJECT, DEFAULT_SUBJECT, "subject")
	bodyPtr              = flag.String(PARAMETER_BODY, DEFAULT_BODY, "body")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := do(
		*smtpHostPtr,
		*smtpPortPtr,
		*smtpTlsPtr,
		*smtpTlsSkipVerifyPtr,
		*fromPtr,
		*toPtr,
		*subjectPtr,
		*bodyPtr,
	)
	if err != nil {
		glog.Exit(err)
	}
}

func do(
	smtpHost string,
	smtpPort int,
	smtpTls bool,
	smtpTlsSkipVerify bool,
	from string,
	to string,
	subject string,
	body string,
) error {
	config := config.New()
	config.SetSmtpHost(smtpHost)
	config.SetSmtpPort(smtpPort)
	config.SetTls(smtpTls)
	config.SetTlsSkipVerify(smtpTlsSkipVerify)
	mailer := mailer.New(config)
	message := message.New()
	message.SetSender(from)
	message.SetRecipient(to)
	message.SetSubject(subject)
	message.SetContent(body)
	if err := mailer.Send(message); err != nil {
		return err
	}
	glog.V(2).Infof("send mail successful")
	return nil
}
