package config

import "github.com/kelseyhightower/envconfig"

// AppConfig contains the app wide configuration
type AppConfig struct {
	// ImapHost is the host address for the IMAP server
	ImapHost HostDecoder `required:"true" split_words:"true"`
	// ImapTlsEnabled is true in case the connection accepts TLS connections or false if it doesn't.
	ImapTlsEnabled bool `default:"true" split_words:"true"`
	// ImapUsername represents the username for the IMAP server
	ImapUsername string `required:"true" split_words:"true"`
	// ImapPassword represents the password set for the user
	ImapPassword string `required:"true" split_words:"true"`
	// NotificationUris contains a list of services that should be notified for mails
	NotificationUris []string `required:"true" split_words:"true"`
}

var config AppConfig

// Read from env variables
func Read() error {
	c := AppConfig{}
	err := envconfig.Process("imap_to_chat_bridge_", &c)
	config = c
	return err
}

// PrintUsage displays the help for the env vars
func PrintUsage() {
	_ = envconfig.Usage("imap_to_chat_bridge_", &AppConfig{})
}

// Get current app configuration, make sure Read has been called before.
func Get() *AppConfig {
	return &config
}
