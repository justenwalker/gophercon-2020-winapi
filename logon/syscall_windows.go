package main

import "strings"

// BOOL LogonUserA(
// 	LPCWSTR  lpszUsername,
// 	LPCWSTR  lpszDomain,
// 	LPCWSTR  lpszPassword,
// 	DWORD   dwLogonType,
// 	DWORD   dwLogonProvider,
// 	PHANDLE phToken
// );
//
//sys _LogonUser(username *uint16, domain *uint16, password *uint16, logonType uint32, logonProvider uint32, token *windows.Token) (err error) = advapi32.LogonUserW

// Logon types
const (
	_LOGON32_LOGON_INTERACTIVE       uint32 = 2
	_LOGON32_LOGON_NETWORK           uint32 = 3
	_LOGON32_LOGON_BATCH             uint32 = 4
	_LOGON32_LOGON_SERVICE           uint32 = 5
	_LOGON32_LOGON_UNLOCK            uint32 = 7
	_LOGON32_LOGON_NETWORK_CLEARTEXT uint32 = 8
	_LOGON32_LOGON_NEW_CREDENTIALS   uint32 = 9
)

// Logon providers
const (
	_LOGON32_PROVIDER_DEFAULT uint32 = 0
	_LOGON32_PROVIDER_WINNT40 uint32 = 2
	_LOGON32_PROVIDER_WINNT50 uint32 = 3
)

//go:generate go run golang.org/x/tools/cmd/stringer  -type=SidType
type SidType uint32

const (
	SidTypeUser SidType = iota + 1
	SidTypeGroup
	SidTypeDomain
	SidTypeAlias
	SidTypeWellKnownGroup
	SidTypeDeletedAccount
	SidTypeInvalid
	SidTypeUnknown
	SidTypeComputer
	SidTypeLabel
	SidTypeLogonSession
)

func (s SidType) Name() string {
	return strings.TrimPrefix(s.String(), "SidType")
}
