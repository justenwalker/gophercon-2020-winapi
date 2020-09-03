package main

import (
	"fmt"
	"time"
	"unsafe"
)

func CredEnumerate(filter string) (*Credentials, error) {
	var cred Credentials
	var err error
	if filter != "" {
		var wstr *uint16
		wstr, err = UTF16PtrFromString(filter)
		if err != nil {
			return nil, err
		}
		err = CredEnumerateW(wstr, 0, &cred.count, &cred.items)
	} else {
		err = CredEnumerateW(nil, _CRED_ENUMERATE_ALL_CREDENTIALS, &cred.count, &cred.items)
	}
	if err != nil {
		return nil, err
	}
	return &cred, nil
}

type Credentials struct {
	free  bool
	count uint32
	items **_CREDENTIALW
}

func (c *Credentials) ForEach(fn func(cred Credential) error) error {
	if c.free {
		return fmt.Errorf("memory freed")
	}
	var sz = unsafe.Sizeof(&_CREDENTIALW{})
	for i := uint32(0); i < c.count; i++ {
		pcred := *(**_CREDENTIALW)(unsafe.Pointer(uintptr(unsafe.Pointer(c.items)) + uintptr(i)*sz))
		if err := fn(Credential{
			free:  &c.free,
			pcred: pcred,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (c *Credentials) Free() {
	if !c.free {
		CredFree(unsafe.Pointer(c.items))
		c.free = true
	}
}

type Credential struct {
	free  *bool
	pcred *_CREDENTIALW
}

func (c *Credential) Name() string {
	if *c.free {
		return ""
	}
	return UTF16PtrToString(c.pcred.TargetName)
}

func (c *Credential) Alias() string {
	if *c.free {
		return ""
	}
	return UTF16PtrToString(c.pcred.TargetAlias)
}

func (c *Credential) Comment() string {
	if *c.free {
		return ""
	}
	return UTF16PtrToString(c.pcred.Comment)
}

func (c *Credential) UserName() string {
	if *c.free {
		return ""
	}
	return UTF16PtrToString(c.pcred.UserName)
}

func (c *Credential) Credential() []byte {
	if *c.free {
		return nil
	}
	if c.pcred.CredentialBlobSize == 0 {
		return nil
	}
	n := int(c.pcred.CredentialBlobSize)
	return (*[1 << 30]byte)(unsafe.Pointer(c.pcred.CredentialBlob))[0:n:n]
}

func (c *Credential) LastWritten() time.Time {
	if *c.free {
		return time.Time{}
	}
	return c.pcred.LastWritten.ToTime()
}

func (c *Credential) Type() string {
	if *c.free {
		return ""
	}
	switch c.pcred.Type {
	case _CRED_TYPE_DOMAIN_CERTIFICATE:
		return "Domain Cert"
	case _CRED_TYPE_DOMAIN_EXTENDED:
		return "Domain Extended"
	case _CRED_TYPE_DOMAIN_PASSWORD:
		return "Domain Password"
	case _CRED_TYPE_DOMAIN_VISIBLE_PASSWORD:
		return "Domain Visible Password"
	case _CRED_TYPE_GENERIC:
		return "Generic"
	case _CRED_TYPE_GENERIC_CERTIFICATE:
		return "Generic Certificate"
	default:
		return "Unknown"
	}
}

type CredentialAttribute struct {
	Keyword string
	Flags   uint32
	Value   []byte
}

func (c *Credential) Attributes() []CredentialAttribute {
	if *c.free {
		return nil
	}
	if c.pcred.AttributeCount == 0 {
		return nil
	}
	attrs := make([]CredentialAttribute, 0, int(c.pcred.AttributeCount))
	var sz = unsafe.Sizeof(_CREDENTIAL_ATTRIBUTEW{})
	for i := uint32(0); i < c.pcred.AttributeCount; i++ {
		pattr := (*_CREDENTIAL_ATTRIBUTEW)(unsafe.Pointer(uintptr(unsafe.Pointer(c.pcred.Attributes)) + uintptr(i)*sz))
		key := UTF16PtrToString(pattr.Keyword)
		n := int(pattr.ValueSize)
		var val []byte
		if n > 0 {
			val = make([]byte, n)
			copy(val, (*(*[256]byte)(unsafe.Pointer(pattr.Value)))[:n:n])
		}
		attrs = append(attrs, CredentialAttribute{
			Keyword: key,
			Flags:   pattr.Flags,
			Value:   val,
		})
	}
	return attrs
}
