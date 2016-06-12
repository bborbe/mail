package config

type config struct {
	smtpUser          string
	smtpPassword      string
	smtpHost          string
	smtpPort          int
	smtpTls           bool
	smtpTlsSkipVerify bool
}

func New() *config {
	c := new(config)
	c.smtpTls = true
	c.smtpTlsSkipVerify = false
	c.smtpPort = 25
	return c
}

func (m *config) SmtpUser() string {
	return m.smtpUser
}

func (m *config) SetSmtpUser(smtpUser string) {
	m.smtpUser = smtpUser
}

func (m *config) SmtpPassword() string {
	return m.smtpPassword
}

func (m *config) SetSmtpPassword(smtpPassword string) {
	m.smtpPassword = smtpPassword
}

func (m *config) SmtpHost() string {
	return m.smtpHost
}

func (m *config) SetSmtpHost(smtpHost string) {
	m.smtpHost = smtpHost
}

func (m *config) SmtpPort() int {
	return m.smtpPort
}

func (m *config) SetSmtpPort(smtpPort int) {
	m.smtpPort = smtpPort
}

func (m *config) Tls() bool {
	return m.smtpTls
}

func (m *config) SetTls(tls bool) {
	m.smtpTls = tls
}

func (m *config) TlsSkipVerify() bool {
	return m.smtpTlsSkipVerify
}

func (m *config) SetTlsSkipVerify(tlsSkipVerify bool) {
	m.smtpTlsSkipVerify = tlsSkipVerify
}
