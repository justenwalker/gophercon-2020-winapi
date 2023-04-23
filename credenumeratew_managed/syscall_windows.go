package main

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

//sys CredEnumerateW(filter *uint16, flags uint32, count *uint32, credentials ***_CREDENTIALW) (err error) [failretval==0] = advapi32.CredEnumerateW
//sys CredFree(buffer unsafe.Pointer)  = advapi32.CredFree
//sys FileTimeToSystemTime(fileTime *_FILETIME, systemTime *_SYSTEMTIME) [failretval==0] (err error) = kernel32.FileTimeToSystemTime

// UTF16PtrFromString converts a string to a UTF-16 C-String
func UTF16PtrFromString(str string) (*uint16, error) {
	return syscall.UTF16PtrFromString(str)
}

const wcharSize = uintptr(2)

// UTF16PtrToString is like syscall.UTF16ToString, but takes *uint16
// as a parameter instead of []uint16.
func UTF16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	end := unsafe.Pointer(p)
	var n int
	for *(*uint16)(end) != 0 { // Advance to the NULL terminator
		// end = unsafe.Pointer(uintptr(end) + wcharSize)
		end = unsafe.Add(end, wcharSize)
		n++
	}
	// Convert *uint16 to []uint16
	wstr := unsafe.Slice(p, n)
	return string(utf16.Decode(wstr))
}
