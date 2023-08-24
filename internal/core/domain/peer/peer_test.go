package peer

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalPeers(t *testing.T) {
	tests := map[string]struct {
		input  string
		output struct {
			peers []Peer
			err   error
		}
	}{
		"correctly parses peers": {
			input: string([]byte{127, 0, 0, 1, 0x00, 0x50, 1, 1, 1, 1, 0x01, 0xbb, 192, 168, 0, 0, 0x1f, 0x90}),
			output: struct {
				peers []Peer
				err   error
			}{
				peers: []Peer{
					{IP: net.IP{127, 0, 0, 1}, Port: 80},
					{IP: net.IP{1, 1, 1, 1}, Port: 443},
					{IP: net.IP{192, 168, 0, 0}, Port: 8080},
				},
				err: nil,
			},
		},
		"not enough bytes in peers": {
			input: string([]byte{127, 0, 0, 1, 0x00}),
			output: struct {
				peers []Peer
				err   error
			}{
				peers: nil,
				err:   ErrMalFormedPeers,
			},
		},
	}

	for _, test := range tests {
		peers, err := UnmarshalPeers([]byte(test.input))
		assert.Equal(t, test.output.peers, peers)
		assert.Equal(t, test.output.err, err)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		input  Peer
		output string
	}{
		{
			input:  Peer{IP: net.IP{127, 0, 0, 1}, Port: 8080},
			output: "127.0.0.1:8080",
		},
	}
	for _, test := range tests {
		s := test.input.String()
		assert.Equal(t, test.output, s)
	}
}
