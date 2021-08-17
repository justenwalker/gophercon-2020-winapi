package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

func main() {
	var domain string
	username, err := readLine("Username", false)
	die("failed to read username", err)
	password, err := readLine("Password", true)
	die("failed to read password", err)
	fmt.Println("Checking Password")
	if i := strings.IndexByte(username, '\\'); i != -1 {
		username, domain = string(username[0:i]), string(username[i+1:])
	}
	user, err := LogonUser(username, domain, password)
	die("Login failed", err)
	fmt.Printf("Welcome, %s\n", user.Name)
	js, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("%s\n", string(js))
}

func die(msg string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Login failed:", err)
		os.Exit(1)
	}
}

func readLine(label string, secret bool) (string, error) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
	}
	if secret {
		prompt.Mask = '*'
	}
	result, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}
	return result, nil
}
