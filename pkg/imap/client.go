package imap

import (
	"crypto/tls"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"jaytaylor.com/html2text"
	"strconv"
)

type emailResult struct {
	Email *EMail
	Error error
}

func emailResultFromError(err error) emailResult {
	return emailResult{Error: err}
}

func newImapClient(host string, port int, useTls bool) (*client.Client, error) {
	address := host + ":" + strconv.Itoa(port)
	if useTls {
		return client.DialTLS(address, &tls.Config{})
	}
	return client.Dial(address)
}

// Client is a high level API for IMAP
type Client struct {
	imap *client.Client
}

// NewClient creates a new TCP connection and logs the user in
func NewClient(host string, port int, useTls bool, mail string, password string) (*Client, error) {
	imapClient, err := newImapClient(host, port, useTls)
	if err != nil {
		return nil, err
	}

	err = imapClient.Login(mail, password)
	if err != nil {
		return nil, err
	}

	return &Client{imap: imapClient}, nil
}

// DeleteMail by uid, marks for deletion
func (c *Client) DeleteMail(uid uint32) error {
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(uid)
	err := c.imap.UidStore(seqSet, "+FLAGS", []interface{}{imap.DeletedFlag}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetMails for a given folder from the server
func (c *Client) GetMails(folder string) chan emailResult {
	results := make(chan emailResult)
	_, err := c.imap.Select(folder, false)
	if err != nil {
		results <- emailResultFromError(err)
		close(results)
		return results
	}

	criteria := imap.NewSearchCriteria()
	messageIds, err := c.imap.Search(criteria)
	if err != nil {
		results <- emailResultFromError(err)
		close(results)
		return results
	}

	if len(messageIds) == 0 {
		close(results)
		return results
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(messageIds...)
	messages := make(chan *imap.Message, len(messageIds))
	done := make(chan error, 1)
	bodySection := &imap.BodySectionName{}
	go func() {
		done <- c.imap.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope, imap.FetchBodyStructure, bodySection.FetchItem()}, messages)
	}()

	go func() {
		for msg := range messages {
			email, err := c.parseEmail(msg, bodySection)
			if err != nil {
				results <- emailResult{email, err}
				continue
			}

			results <- emailResult{Email: email}

			if err := <-done; err != nil {
				results <- emailResultFromError(err)
			}
		}

		close(results)
	}()

	return results
}

func (c *Client) parseEmail(msg *imap.Message, bodySection *imap.BodySectionName) (*EMail, error) {
	r := msg.GetBody(bodySection)
	mr, err := mail.CreateReader(r)
	if err != nil {
		return nil, err
	}

	email := EMail{
		Uid:         msg.Uid,
		Subject:     msg.Envelope.Subject,
		From:        msg.Envelope.From,
		Attachments: make(map[string][]byte),
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			// This is the message's text (can be plain-text or HTML)
			b, _ := io.ReadAll(p.Body)
			ctype, _, _ := h.ContentType()
			switch ctype {
			case "text/html":
				email.Html = string(b)
				break
			case "text/plain":
			case "text":
				email.PlainText = string(b)
				break
			default:
				break
			}
		case *mail.AttachmentHeader:
			filename, _ := h.Filename()
			b, _ := io.ReadAll(p.Body)
			email.Attachments[filename] = b
		}
	}

	if email.PlainText == "" {
		text, err := html2text.FromString(email.Html)
		if err != nil {
			return nil, err
		}
		email.PlainText = text
	}

	return &email, nil
}

// Close connection and delete all messages marked as delete
func (c *Client) Close() {
	_ = c.imap.Close()
}
