package config

type config struct {
	smtpUser     string
	smtpPassword string
	smtpHost     string
	smtpPort     int
}

func New() *config {
	return new(config)
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
