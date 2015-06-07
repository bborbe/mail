package mail

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsMail(t *testing.T) {
	c := New(nil)
	var i *Mailer
	err := AssertThat(c, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}
