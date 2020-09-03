package main

import (
	"fmt"
	"log"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

func main() {
	rows, err := GetTCPTable()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%-4s\t%-21s\t%-21s\t%s\n","PID","LOCAL","REMOTE","STATE")
	for _, row := range rows {
		remote := row.Remote.String()
		if remote == ":0" {
			remote = ""
		}
		fmt.Printf("% 4d\t%-21s\t%-21s\t%s\n", row.PID, row.Local, remote, row.State)
	}
}