package message

type message struct {
	sender    string
	recipient string
	content   string
}

func New() *message {
	return new(message)
}

func (m *message) Sender() string {
	return m.sender
}
func (m *message) Recipient() string {
	return m.recipient
}

func (m *message) Content() string {
	return m.content
}

func (m *message) SetSender(sender string) {
	m.sender = sender
}
func (m *message) SetRecipient(recipient string) {
	m.recipient = recipient
}

func (m *message) SetContent(content string) {
	m.content = content
}
