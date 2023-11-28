package config

import (
	"net"
	"strconv"
)

// HostInfo represents the parsed hostname and port of a host address
type HostInfo struct {
	// Hostname represents the FQDN or IP of the host
	Hostname string
	// Port represents the port to connect to
	Port int
}

// HostDecoder converts the given string value into host and port pair
type HostDecoder HostInfo

// Decode string into hostname and port
func (hd *HostDecoder) Decode(value string) error {
	host, port, err := net.SplitHostPort(value)
	if err != nil {
		return err
	}
	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	*hd = HostDecoder(HostInfo{
		host,
		parsedPort,
	})

	return err
}
