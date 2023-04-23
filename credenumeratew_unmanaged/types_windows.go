package main

import (
	"time"
)

const _CRED_ENUMERATE_ALL_CREDENTIALS uint32 = 0x1

/*
	typedef struct _FILETIME {
		DWORD dwLowDateTime;
		DWORD dwHighDateTime;
	} FILETIME, *PFILETIME, *LPFILETIME;
*/
type _FILETIME struct {
	LowDateTime  uint32
	HighDateTime uint32
}

func (t *_FILETIME) ToTime() time.Time {
	var systime _SYSTEMTIME
	if err := FileTimeToSystemTime(t, &systime); err != nil {
		return time.Time{}
	}
	return systime.ToTime()
}

/*
	typedef struct _SYSTEMTIME {
	  WORD wYear;
	  WORD wMonth;
	  WORD wDayOfWeek;
	  WORD wDay;
	  WORD wHour;
	  WORD wMinute;
	  WORD wSecond;
	  WORD wMilliseconds;
	} SYSTEMTIME, *PSYSTEMTIME, *LPSYSTEMTIME;
*/
type _SYSTEMTIME struct {
	Year         uint16
	Month        uint16
	DayOfWeek    uint16
	Day          uint16
	Hour         uint16
	Minute       uint16
	Second       uint16
	Milliseconds uint16
}

func (s *_SYSTEMTIME) ToTime() time.Time {
	return time.Date(
		int(s.Year),
		time.Month(s.Month),
		int(s.Day),
		int(s.Hour),
		int(s.Minute),
		int(s.Second),
		int(s.Milliseconds)*1_000_000, time.UTC)
}

/*
	typedef struct _CREDENTIALW {
	  DWORD                  Flags;
	  DWORD                  Type;
	  LPWSTR                 TargetName;
	  LPWSTR                 Comment;
	  FILETIME               LastWritten;
	  DWORD                  CredentialBlobSize;
	  LPBYTE                 CredentialBlob;
	  DWORD                  Persist;
	  DWORD                  AttributeCount;
	  PCREDENTIAL_ATTRIBUTEW Attributes;

#if ...

	wchar_t                *TargetAlias;

#else

	LPWSTR                 TargetAlias;

#endif
#if ...

	wchar_t                *UserName;

#else

	LPWSTR                 UserName;

#endif
} CREDENTIALW, *PCREDENTIALW;
*/
type _CREDENTIALW struct {
	Flags              uint32
	Type               uint32
	TargetName         *uint16
	Comment            *uint16
	LastWritten        _FILETIME
	CredentialBlobSize uint32
	CredentialBlob     *byte
	Persist            uint32
	AttributeCount     uint32
	Attributes         *_CREDENTIAL_ATTRIBUTEW
	TargetAlias        *uint16
	UserName           *uint16
}

const (
	_CRED_TYPE_GENERIC                 uint32 = 0x1
	_CRED_TYPE_DOMAIN_PASSWORD         uint32 = 0x2
	_CRED_TYPE_DOMAIN_CERTIFICATE      uint32 = 0x3
	_CRED_TYPE_DOMAIN_VISIBLE_PASSWORD uint32 = 0x4
	_CRED_TYPE_GENERIC_CERTIFICATE     uint32 = 0x5
	_CRED_TYPE_DOMAIN_EXTENDED         uint32 = 0x6
)

/*
	typedef struct _CREDENTIAL_ATTRIBUTEW {
	  LPWSTR  Keyword;
	  DWORD   Flags;
	  DWORD   ValueSize;
	  LPBYTE  Value;
	} CREDENTIAL_ATTRIBUTEW, *PCREDENTIAL_ATTRIBUTEW;
*/
type _CREDENTIAL_ATTRIBUTEW struct {
	Keyword   *uint16
	Flags     uint32
	ValueSize uint32
	Value     *byte
}
