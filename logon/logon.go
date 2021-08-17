package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

type User struct {
	SID    string  `json:"sid"`
	Name   string  `json:"name"`
	Domain string  `json:"domain,omitempty"`
	Type   string  `json:"type"`
	Groups []Group `json:"groups,omitempty"`
}

type Group struct {
	SID    string `json:"sid"`
	Name   string `json:"name"`
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type"`
}

func LogonUser(username, domain, password string) (*User, error) {
	var (
		d *uint16
		u *uint16
		p *uint16
	)
	if domain != "" {
		d, _ = windows.UTF16PtrFromString(domain)
	}
	if username != "" {
		u, _ = windows.UTF16PtrFromString(username)
	}
	if password != "" {
		p, _ = windows.UTF16PtrFromString(password)
	}
	token, err := lookupUser(u, d, p)
	if err != nil {
		return nil, err
	}
	defer token.Close()
	var user User
	tokenuser, err := token.GetTokenUser()
	if err != nil {
		return nil, fmt.Errorf("GetTokenUser: %w", err)
	}
	acc, domain, use, err := tokenuser.User.Sid.LookupAccount("")
	if err != nil {
		return nil, fmt.Errorf("LookupAccount(user=%s): %w", tokenuser.User.Sid.String(), err)
	}
	user.SID = tokenuser.User.Sid.String()
	user.Name = acc
	user.Type = SidType(use).Name()
	user.Domain = domain

	tokengroups, err := token.GetTokenGroups()
	if err != nil {
		return nil, fmt.Errorf("GetTokenGroups: %w", err)
	}
	for _, tokengroup := range tokengroups.AllGroups() {
		acc, domain, use, err := tokengroup.Sid.LookupAccount("")
		if err != nil {
			return nil, fmt.Errorf("LookupAccount(group=%s): %w", tokengroup.Sid.String(), err)
		}
		user.Groups = append(user.Groups, Group{
			SID:    tokengroup.Sid.String(),
			Name:   acc,
			Domain: domain,
			Type:   SidType(use).Name(),
		})
	}
	return &user, nil
}

func lookupUser(u *uint16, d *uint16, p *uint16) (*windows.Token, error) {
	var token windows.Token
	err := _LogonUser(u, d, p, _LOGON32_LOGON_NETWORK, _LOGON32_PROVIDER_DEFAULT, &token)
	if err == nil {
		// default lookup succeeded
		return &token, nil
	}
	// if not, try NEW_CREDENTIALS instead
	if err := _LogonUser(u, d, p, _LOGON32_LOGON_NEW_CREDENTIALS, _LOGON32_PROVIDER_WINNT50, &token); err == nil {
		return &token, nil
	}
	return nil, err
}
