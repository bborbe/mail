package mock

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/mail"
)

func TestImplementsMail(t *testing.T) {
	c := New()
	var i *mail.Mailer
	err := AssertThat(c, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
