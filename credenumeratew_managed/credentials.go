package main

import (
	"time"
	"unsafe"
)

type Credential struct {
	Name        string
	Alias       string
	Type        string
	Comment     string
	UserName    string
	Credential  []byte
	LastWritten time.Time
	Attributes  []CredentialAttribute
}

type CredentialAttribute struct {
	Keyword string
	Flags   uint32
	Value   []byte
}

func CredEnumerate(filter string) ([]Credential, error) {
	var creds []Credential
	var err error
	var count uint32
	var items **_CREDENTIALW
	if filter != "" {
		var wstr *uint16
		wstr, err = UTF16PtrFromString(filter)
		if err != nil {
			return nil, err
		}
		err = CredEnumerateW(wstr, 0, &count, &items)
	} else {
		err = CredEnumerateW(nil, _CRED_ENUMERATE_ALL_CREDENTIALS, &count, &items)
	}
	if err != nil {
		return nil, err
	}
	// Free Immediately before returning
	defer CredFree(unsafe.Pointer(items))

	var sz = unsafe.Sizeof(&_CREDENTIALW{})
	creds = make([]Credential, 0, int(count))
	for i := uint32(0); i < count; i++ {
		pcred := *(**_CREDENTIALW)(unsafe.Pointer(uintptr(unsafe.Pointer(items)) + uintptr(i)*sz))
		// Copy all data to Go structs in Managed Memory
		creds = append(creds, toCredential(pcred))
	}
	return creds, nil
}

func toCredential(pcred *_CREDENTIALW) (credential Credential) {
	credential.Name = UTF16PtrToString(pcred.TargetName)
	credential.Alias = UTF16PtrToString(pcred.TargetAlias)
	credential.Comment = UTF16PtrToString(pcred.Comment)
	credential.UserName = UTF16PtrToString(pcred.UserName)
	credential.LastWritten = pcred.LastWritten.ToTime()
	switch pcred.Type {
	case _CRED_TYPE_DOMAIN_CERTIFICATE:
		credential.Type = "Domain Cert"
	case _CRED_TYPE_DOMAIN_EXTENDED:
		credential.Type = "Domain Extended"
	case _CRED_TYPE_DOMAIN_PASSWORD:
		credential.Type = "Domain Password"
	case _CRED_TYPE_DOMAIN_VISIBLE_PASSWORD:
		credential.Type = "Domain Visible Password"
	case _CRED_TYPE_GENERIC:
		credential.Type = "Generic"
	case _CRED_TYPE_GENERIC_CERTIFICATE:
		credential.Type = "Generic Certificate"
	default:
		credential.Type = "Unknown"
	}
	if pcred.CredentialBlobSize > 0 {
		n := int(pcred.CredentialBlobSize)
		val := make([]byte, n)
		copy(val, unsafe.Slice(pcred.CredentialBlob, n))
		credential.Credential = val
	}
	if pcred.AttributeCount > 0 {
		credential.Attributes = make([]CredentialAttribute, 0, int(pcred.AttributeCount))
		var attrsz uintptr = unsafe.Sizeof(_CREDENTIAL_ATTRIBUTEW{})
		for i := uint32(0); i < pcred.AttributeCount; i++ {
			pattr := (*_CREDENTIAL_ATTRIBUTEW)(unsafe.Pointer(uintptr(unsafe.Pointer(pcred.Attributes)) + uintptr(i)*attrsz))
			key := UTF16PtrToString(pattr.Keyword)
			n := int(pattr.ValueSize)
			var val []byte
			if n > 0 {
				val = make([]byte, n)
				copy(val, (*(*[256]byte)(unsafe.Pointer(pattr.Value)))[:n:n])
			}
			credential.Attributes = append(credential.Attributes, CredentialAttribute{
				Keyword: key,
				Flags:   pattr.Flags,
				Value:   val,
			})
		}
	}
	return credential
}
