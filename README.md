# Safety Not Guaranteed - GopherCon 2020

This repo contains some full-featured examples of calling Windows APIs from Go from my talk at GopherCon.
Check out each folder for mor information.

## Building and Running the Examples

Run `.\build.ps1`, the commands will be built into the `.\bin` folder

## Unsafe Operations

Here is a list of collection unsafe operations that should be in your tool belt.

### Pointer Arithmetic

```go
var arr *T // Pointer to the first element of the array
i := uintptr(1) // subscript/offset to the element we want
sz := unsafe.Sizeof(T{}) // sz = size of T in bytes

// t := arr[i]
t := *(*T)unsafe.Pointer(uintptr(unsafe.Pointer(arr)) + (i*sz))
// this is several steps all happening on one line:
//  a) uintptr(unsafe.Pointer(arr)) --- convert to uintptr so we can do arithmetic
//  b) (a + i*sz) - advance the pointer the correct number of bytes to get to index 'i'
//  c) (*T)unsafe.Pointer(b) -- convert uintptr back to a *T
//  d) *(c) -- dereference *T to get the value T
```

*Note*: As of Go 1.17, the unsafe package has [`unsafe.Add`](https://pkg.go.dev/unsafe#Add) making pointer-arithmetic more straight-forward.

```go
var arr *T // Pointer to the first element of the array
i := uintptr(1) // subscript/offset to the element we want
sz := unsafe.Sizeof(T{}) // sz = size of T in bytes

// t := arr[i]
t := *(*T)unsafe.Add(unsafe.Pointer(arr),i * sz)
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
		// Go 1.16 and earlier
		p = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1))
		// Go 1.17 and later
		// p = (*uint8)unsafe.Add(unsafe.Pointer(p),1)
		n++
	}
    return
}
```

### Converting from *T,len => []T

A C array is a pointer to the first element. 
The size of the array is either given, or found (rather unsafely) by traversing the array until encountering a sentinel (usually `NULL`).

```go
var tptr *T // Pointer to the first element of the array
var n int   // The real length of the array

// Convert *T,n to []T
ts := (*[1 << 30]T)(unsafe.Pointer(tptr))[:n:n]
// this is 3 steps all happening at once:
//  a) unsafe.Pointer(tptr) --- so we can convert to a pointer of another type
//  b) (*[1<<30]T)(a) -- convert to a pointer to a large array.
//  c) (b)[0:n:n] -- create a slice backed by the array we're pointing at, setting both its length and capacity to the known value 'n'
```

*Note*: As of Go 1.17, the unsafe package has [`unsafe.Slice(ptr *T, len anyIntegerType) []T`](https://pkg.go.dev/unsafe#Slice) making this construct obsolete 🎉.

```go
var tptr *T // Pointer to the first element of the array
var n int   // the real length of the array

// Convert *T,n to []T
ts := unsafe.Slice(tptr,n)
```