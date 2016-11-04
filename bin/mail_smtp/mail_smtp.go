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
	defaultHost            = "localhost"
	defaultPort            = 1025
	defaultTls             = false
	defaultTlsSkipVerify   = false
	defaultFrom            = "test@example.com"
	defaultTo              = "test@example.com"
	defaultBody            = "Hello World\r\n"
	defaultSubject         = "Test Mail"
	parameterSmtpHost      = "smtp-host"
	parameterSmtpPort      = "smtp-port"
	parameterTls           = "smtp-tls"
	parameterTlsSkipVerify = "smtp-tls-skip-verify"
	parameterFrom          = "from"
	parameterTo            = "to"
	parameterSubject       = "subject"
	parameterBody          = "body"
)

var (
	smtpHostPtr          = flag.String(parameterSmtpHost, defaultHost, "smtp host")
	smtpPortPtr          = flag.Int(parameterSmtpPort, defaultPort, "smtp port")
	smtpTlsPtr           = flag.Bool(parameterTls, defaultTls, "smtp tls")
	smtpTlsSkipVerifyPtr = flag.Bool(parameterTlsSkipVerify, defaultTlsSkipVerify, "smtp tls skip verify")
	fromPtr              = flag.String(parameterFrom, defaultFrom, "from")
	toPtr                = flag.String(parameterTo, defaultTo, "to")
	subjectPtr           = flag.String(parameterSubject, defaultSubject, "subject")
	bodyPtr              = flag.String(parameterBody, defaultBody, "body")
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
