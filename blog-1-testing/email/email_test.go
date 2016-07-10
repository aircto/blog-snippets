package email

import "testing"

type FakeEmailSender struct {
	to   []string
	from string
	body []byte
}

func (f *FakeEmailSender) SendEmail(to []string, from, subject string, body []byte) error {
	f.to, f.from, f.body = to, from, body
	return nil
}

func init() {
	Sender = &FakeEmailSender{}
}

func TestSendEmailToAdmin(t *testing.T) {

}
