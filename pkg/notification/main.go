package notification

import (
	"github.com/containrrr/shoutrrr"
	"github.com/timo-reymann/imap-to-chat-bridge/pkg/imap"
)

// SendToChat delivers a message using shouttr, delivered from the mail provided
func SendToChat(url string, mail *imap.EMail) error {
	return shoutrrr.Send(url, mail.Subject+"\n\n"+mail.PlainText)
}
