# Safety Not Guaranteed - GopherCon 2020

This repo contains some full-featured examples of calling Windows APIs from Go from my talk at GopherCon.
Check out each folder for mor information.

## Building and Running the Examples

1. Run `go get -u golang.org/x/sys/windows` to make sure you have it in your `GOPATH`.
2. Run `go build` in the example folders to build the executable.
3. Run the `.exe` to test the program.

*Note*: I've only tested this on Go 1.15. I see no reason why this wouldn't work on 1.14 though.

## Unsafe Operations

Here is a list of collection unsafe operations that should be in your tool belt.

### Pointer Arithmetic

```go
// var arr *T
i := uintptr(1) // i = subscript/offset to the element we want
sz := unsafe.Sizeof(T{}) // sz = size of T in bytes
t := *(*T)unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + (i*sz))
// this is several steps all happening on one line:
//  a) uintptr(unsafe.Pointer(arr)) --- convert to uintptr so we can do arithmetic
//  b) (a + i*sz) - advance the pointer the correct number of bytes to get to index 'i'
//  c) (*T)unsafe.Pointer(b) -- convert uintptr back to a *T
//  d) *(c) -- dereference *T to get the value T
```

Example:

C-Strings are either *uint8 (ANSI) or *uint16 (UTF-16)
But regardless, to find the length, use pointer arithmetic.

```go
func strlen(p *uint8) (n int) {
	if p == nil {
		return
	}
	for *p != 0 {
		p = unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)
		n++
	}
    return
}
```

*Note*: In a future version of Go, the unsafe package may have [`unsafe.Add`](https://github.com/golang/go/issues/40481) making pointer-arithmetic more straight-forward.

### Converting from *T,len => []T

A C array is a pointer to the first element. 
The size of the array is either given, or found (rather unsafely) by traversing the array until encountering a sentinel (usually `NULL`).

```go
// n provided as the real length of the array
ts := (*[1 << 30]T)(unsafe.Pointer(tptr))[:n:n]
// this is 3 steps all happening at once:
//  a) unsafe.Pointer(tptr) --- so we can convert to a pointer of another type
//  b) (*[1<<30]T)(a) -- convert to a pointer to a large array.
//  c) (b)[0:n:n] -- create a slice backed by the array we're pointing at, setting both its length and capacity to the known value 'n'
```

*Note*: In a future version of Go, the unsafe package may have [`unsafe.Slice(ptr *T, len anyIntegerType) []T`](https://github.com/golang/go/issues/19367) making this construct obsolete 🎉.