package cmd

import (
	"github.com/timo-reymann/imap-to-chat-bridge/pkg/buildinfo"
	"github.com/timo-reymann/imap-to-chat-bridge/pkg/config"
	"github.com/timo-reymann/imap-to-chat-bridge/pkg/imap"
	"github.com/timo-reymann/imap-to-chat-bridge/pkg/notification"
	"log"
	"time"
)

// Execute cmd entrypoint
func Execute() {
	buildinfo.PrintVersionInfo()
	err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	c := config.Get()

	for {
		client, err := imap.NewClient(c.ImapHost.Hostname, c.ImapHost.Port, c.ImapTlsEnabled, c.ImapUsername, c.ImapPassword)
		if err != nil {
			log.Fatal(err)
		}

		results := client.GetMails("INBOX")
		for result := range results {
			if result.Error != nil {
				log.Println("Error fetching mails: " + result.Error.Error())
				continue
			}

			email := result.Email

			err := notification.SendToChat(config.Get().NotificationUris[0], email)
			if err != nil {
				log.Println("Failed to send mail '" + email.Subject + "' to chat: " + err.Error())
			} else {
				err := client.DeleteMail(email.Uid)
				if err != nil {
					continue
				}
			}
		}

		client.Close()
		println("Polling in 10s again ...")
		time.Sleep(10 * time.Second)
	}

}
