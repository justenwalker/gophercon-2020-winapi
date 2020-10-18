# Unsafe Cast

This is an example of using `unsafe.Pointer` to convert a type between two different forms (like a C union).
This is legal because the two types are the same memory size, so when you dereference them, you are not extending into 
unallocated memory. 

You can always cast to a value where the size is equal to or less than the original size,
but the result may be surprising. 

For example, take note of the values you get when it is run.
Given `t1.a=0xC0DECAFE`, you might expect `t2.a` to be `0xC0DE` and t2.b to be `0xCAFE`,
but instead you get:

```
t1.a=0xC0DECAFE

t2.a=0xCAFE
t2.b=0xC0DE
```

Why?

This is the [Endianness](https://en.wikipedia.org/wiki/Endianness) of the memory sequencing mode of the CPU leaking through. 
The memory for `T1` is actually laid-out physically as "Little Endian", or Lower-order bytes first (Intel x86/amd64 are LE).

```
bytes   0  1  2  3
        FE CA DE 0C
       |--|--|--|--|
       \           /
        \_________/
             |
           t1.a
        0xC0DECAFE
```

So when we access them as T2, those fields correspond to

```
bytes   0  1  2  3
        FE CA DE 0C
       |--|--|--|--|
      /     / \     \
     /_____/   \_____\
        |         |
      t2.a      t2.b
     0xCAFE    0xCODE
```

When printed, they display as "Big Endian" (Higher order bytes first)
which is more natural for a human to read. This mode also happens to be "Network Byte Order", the agreed-upon byte ordering when transmitting data over a network.



