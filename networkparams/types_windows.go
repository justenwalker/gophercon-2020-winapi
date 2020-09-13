package main

import "syscall"

const (
	_ERROR_SUCCESS         syscall.Errno = 0
	_ERROR_BUFFER_OVERFLOW syscall.Errno = 0x6F
)

const (
	_MAX_HOSTNAME_LEN = 128
	_MAX_DOMAIN_LEN   = 128
	_MAX_SCOPE_ID_LEN = 256
)

const (
	_BROADCAST_NODETYPE    uint32 = 0x0001
	_PEER_TO_PEER_NODETYPE uint32 = 0x0002
	_MIXED_NODETYPE        uint32 = 0x0004
	_HYBRID_NODETYPE       uint32 = 0x0008
)

/*
typedef struct {
  char            HostName[MAX_HOSTNAME_LEN + 4];
  char            DomainName[MAX_DOMAIN_NAME_LEN + 4];
  PIP_ADDR_STRING CurrentDnsServer;
  IP_ADDR_STRING  DnsServerList;
  UINT            NodeType;
  char            ScopeId[MAX_SCOPE_ID_LEN + 4];
  UINT            EnableRouting;
  UINT            EnableProxy;
  UINT            EnableDns;
} FIXED_INFO_W2KSP1, *PFIXED_INFO_W2KSP1;
*/
type _FIXED_INFO struct {
	HostName         [_MAX_HOSTNAME_LEN + 4]byte
	DomainName       [_MAX_DOMAIN_LEN + 4]byte
	CurrentDnsServer *_IP_ADDR_STRING
	DnsServerList    _IP_ADDR_STRING
	NodeType         uint32
	ScopeId          [_MAX_SCOPE_ID_LEN + 4]byte
	EnableRouting    uint32
	EnableProxy      uint32
	EnableDNS        uint32
}

/*
typedef struct _IP_ADDR_STRING {
  struct _IP_ADDR_STRING *Next;
  IP_ADDRESS_STRING      IpAddress;
  IP_MASK_STRING         IpMask;
  DWORD                  Context;
} IP_ADDR_STRING, *PIP_ADDR_STRING;
*/
type _IP_ADDR_STRING struct {
	Next      *_IP_ADDR_STRING
	IpAddress _IP_ADDRESS_STRING
	IpMask    _IP_ADDRESS_STRING
	Context   uint32
}

type _IP_ADDRESS_STRING [16]byte
