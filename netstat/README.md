# netstat

A toy netstat re-implementation using [`GetExtendedTCPTable`](https://docs.microsoft.com/en-us/windows/win32/api/iphlpapi/nf-iphlpapi-getextendedtcptable). 

It also defines the relevant API Structs: 

- [`MIB_TCPTABLE_OWNER_PID`](https://docs.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcptable_owner_pid)
- [`MIB_TCPROW_OWNER_PID`](https://docs.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcprow_owner_pid)

This demonstrates a type of API memory exchange where the API call expects to be given a buffer to write into. In this example, we allocated a buffer in Go's managed memory, and the API will write into the buffer
if it is large enough, or it will request more memory if it is too small. We repeat the process of growing the bufer
and calling the API function until it succeeds and, in this case, gives back a TCP connection table.

```
> netstat.exe
t> .\netstat.exe
PID     LOCAL                   REMOTE                  STATE
 14432  127.0.0.1:2015                                  LISTEN
 14432  127.0.0.1:2015          127.0.0.1:54550         ESTABLISHED
 5384   127.0.0.1:49671         127.0.0.1:49672         ESTABLISHED
 5384   127.0.0.1:49672         127.0.0.1:49671         ESTABLISHED

...
```

A slightly simpler example of calling APIs like this can be found in [networkparams](../networkparams).