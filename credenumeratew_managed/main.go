package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

func main() {
	creds, err := CredEnumerate("")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, cred := range creds {
		fmt.Printf("---- %s ---\n", cred.Name)
		if cred.Alias != "" {
			fmt.Printf("Alias:        %q\n", cred.Alias)
		}
		if cred.Comment != "" {
			fmt.Printf("Comment:      %q\n", cred.Comment)
		}
		fmt.Printf("Type:         %s\n", cred.Type)
		if cred.UserName != "" {
			fmt.Printf("UserName:     %q\n", cred.UserName)
		}
		if len(cred.Credential) > 0 {
			hash := sha256.Sum256(cred.Credential)
			masked := hex.EncodeToString(hash[:])[:8]
			fmt.Printf("Cred(masked):  %s\n", masked)
		}
		if len(cred.Attributes) > 0 {
			fmt.Println("Attributes:")
			for _, attr := range cred.Attributes {
				fmt.Printf("\t%s (flags=%#04x): %d bytes\n\t\t%s\n", attr.Keyword, attr.Flags, len(attr.Value), base64.StdEncoding.EncodeToString(attr.Value))
			}
		}
	}
}
