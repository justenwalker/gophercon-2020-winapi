package main

import (
	"net"
	"strconv"
	"unsafe"
)

type NodeType int

const (
	BroadcastNodeType  = NodeType(_BROADCAST_NODETYPE)
	PeerToPeerNodeType = NodeType(_PEER_TO_PEER_NODETYPE)
	MixedNodeType      = NodeType(_MIXED_NODETYPE)
	HybridNodeType     = NodeType(_HYBRID_NODETYPE)
)

func (n NodeType) String() string {
	switch n {
	case BroadcastNodeType:
		return "Broadcast"
	case PeerToPeerNodeType:
		return "Peer to Peer"
	case MixedNodeType:
		return "Mixed"
	case HybridNodeType:
		return "Hybrid"
	}
	return "Unknown " + strconv.FormatUint(uint64(n), 10)
}

type NetworkParams struct {
	HostName      string
	DomainName    string
	DNSServerList []net.IP
	NodeType      NodeType
	ScopeID       string
	EnableRouting bool
	EnableProxy   bool
	EnableDNS     bool
}

func GetNetworkParams() (*NetworkParams, error) {
	bufSize := uint32(unsafe.Sizeof(_FIXED_INFO{}))
	buffer := make([]byte, int(bufSize))
	for {
		ret := _GetNetworkParams(&buffer[0], &bufSize)
		if ret == _ERROR_SUCCESS {
			break
		}
		if ret != _ERROR_BUFFER_OVERFLOW {
			return nil, ret
		}
		buffer = make([]byte, int(bufSize))
	}
	pFixedInfo := (*_FIXED_INFO)(unsafe.Pointer(&buffer[0]))
	return convertToNetworkParams(pFixedInfo)
}

func convertToNetworkParams(pFixedInfo *_FIXED_INFO) (*NetworkParams, error) {
	var np NetworkParams
	np.HostName = sliceToString(pFixedInfo.HostName[:])
	np.DomainName = sliceToString(pFixedInfo.DomainName[:])
	for ipAddr := &pFixedInfo.DnsServerList; ipAddr != nil; ipAddr = ipAddr.Next {
		np.DNSServerList = append(np.DNSServerList, net.ParseIP(sliceToString(ipAddr.IpAddress[:])).To4())
	}
	np.NodeType = NodeType(pFixedInfo.NodeType)
	np.ScopeID = sliceToString(pFixedInfo.ScopeId[:])
	np.EnableRouting = pFixedInfo.EnableRouting != 0
	np.EnableProxy = pFixedInfo.EnableProxy != 0
	np.EnableDNS = pFixedInfo.EnableDNS != 0
	return &np, nil
}

func sliceToString(b []byte) string {
	for i := range b {
		if b[i] == 0 {
			return string(b[0:i])
		}
	}
	return string(b)
}
