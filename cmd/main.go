package cmd

import "github.com/timo-reymann/imap-to-chat-bridge/pkg/buildinfo"

func Execute() {
	buildinfo.PrintVersionInfo()
}
