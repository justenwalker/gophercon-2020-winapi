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
	ce, err := CredEnumerate("")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer ce.Free()
	ce.ForEach(func(cred Credential) error {
		fmt.Printf("---- %s ---\n", cred.Name())
		if alias := cred.Alias(); alias != "" {
			fmt.Printf("Alias:        %q\n", alias)
		}
		if comment := cred.Comment(); comment != "" {
			fmt.Printf("Comment:      %q\n", comment)
		}
		fmt.Printf("Type:         %s\n", cred.Type())
		if username := cred.UserName(); username != "" {
			fmt.Printf("UserName:     %q\n", username)
		}
		if c := cred.Credential(); c != nil {
			hash := sha256.Sum256(c)
			masked := hex.EncodeToString(hash[:])[:8]
			fmt.Printf("Cred(masked):  %s\n", masked)
		}
		if attrs := cred.Attributes(); len(attrs) > 0 {
			fmt.Println("Attributes:")
			for _, attr := range cred.Attributes() {
				fmt.Printf("\t%s (flags=%#04x): %d bytes\n\t\t%s\n", attr.Keyword, attr.Flags, len(attr.Value), base64.StdEncoding.EncodeToString(attr.Value))
			}
		}
		return nil
	})
}
