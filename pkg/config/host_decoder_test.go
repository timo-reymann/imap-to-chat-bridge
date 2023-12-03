package config

import (
	"errors"
	"testing"
)

func TestHostDecoder_Decode(t *testing.T) {
	testCases := []struct {
		input       string
		hostInfo    HostInfo
		expectedErr error
	}{
		{
			"localhost:993",
			HostInfo{"localhost", 993},
			nil,
		},
		{
			"localhost",
			HostInfo{},
			errors.New("address localhost: missing port in address"),
		},
		{
			"localhost:abc",
			HostInfo{},
			errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
		},
	}

	for _, tc := range testCases {
		hd := HostDecoder{}
		err := hd.Decode(tc.input)
		if (tc.expectedErr == nil && err != nil) || (tc.expectedErr != nil && err.Error() != tc.expectedErr.Error()) {
			t.Fatal(err)
		}

		if hd.Port != tc.hostInfo.Port {
			t.Fatalf("Expected port %d, got %d", tc.hostInfo.Port, hd.Port)
		}

		if hd.Hostname != tc.hostInfo.Hostname {
			t.Fatalf("Expected hostname %s, got %s", tc.hostInfo.Hostname, hd.Hostname)
		}
	}
}
