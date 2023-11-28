package imap

import "github.com/emersion/go-imap"

// EMail represents a parsed IMAP message
type EMail struct {
	// Uid on the mail server
	Uid uint32
	// Subject of the mail
	Subject string
	// PlainText representation of the mail. This can be the original one or the stripped down HTML
	PlainText string
	// Html contains the original HTML representation of the mail
	Html string
	// From contains the sender address
	From []*imap.Address
	// Attachments contains a list of attached files
	Attachments map[string][]byte
}
