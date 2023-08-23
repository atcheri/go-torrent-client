package domain

import (
	"encoding/binary"
	"errors"
	"net"
)

var (
	ErrMalFormedPeers = errors.New("received malformed peers")
)

// Peer encodes connection information for a peer
type Peer struct {
	IP   net.IP
	Port uint16
}

// Unmarshal parses peer IP addresses and ports from a buffer
func UnmarshalPeers(peersBin []byte) ([]Peer, error) {
	const peerSize = 6 // 4 for IP, 2 for port
	numPeers := len(peersBin) / peerSize
	if len(peersBin)%peerSize != 0 {
		return nil, ErrMalFormedPeers
	}

	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peer := Peer{
			IP:   net.IP(peersBin[offset : offset+4]),
			Port: binary.BigEndian.Uint16(peersBin[offset+4 : offset+6]),
		}
		peers[i] = peer
	}

	return peers, nil
}
