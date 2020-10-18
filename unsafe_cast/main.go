package main

import (
	"fmt"
	"unsafe"
)

type T1 struct {
	a uint32
}
type T2 struct {
	a uint16
	b uint16
}

func main() {
	t1 := T1{a: 0xC0DECAFE}
	t2 := *(*T2)(unsafe.Pointer(&t1))
	fmt.Printf("sizeof(T1)=%d\n", unsafe.Sizeof(t1))
	fmt.Printf("sizeof(T2)=%d\n", unsafe.Sizeof(t2))
	fmt.Println()
	fmt.Printf("t1.a=0x%X\n\n", t1.a)
	fmt.Printf("t2.a=0x%X\nt2.b=0x%X\n\n", t2.a, t2.b)
	// Output:
	//sizeof(T1)=4
	//sizeof(T2)=4
	//
	//t1.a=0xC0DECAFE
	//
	//t2.a=0xCAFE
	//t2.b=0xC0DE
}