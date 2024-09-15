package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/SabariGanesh-K/21BPS1209_Backend.git/util"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)


	subject := "testing"
	content := `
	<h1>Hello world</h1>
	<p>Hire me pls. Know more about me <a href="http://sabari.codes">www.sabari.codes</a></p>
	`
	to := []string{"k.sabarii.ganesh@gmail.com"}
	attachFiles := []string{"./random.txt"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
