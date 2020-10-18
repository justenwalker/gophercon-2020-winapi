# Network Params

This program lists your computer's network parameters using [`GetNetworkParams`](https://docs.microsoft.com/en-us/windows/win32/api/iphlpapi/nf-iphlpapi-getnetworkparams)

It also defines the relevant API Structs: 

- [`FIXED_INFO`](https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-fixed_info_w2ksp1)
- [`IP_ADDR_STRING `](https://docs.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-ip_addr_string)

The interesting bits are the Host Name, Domain Name (dns search domain), and the list of DNS servers your computer uses for name resolution.
This an example of providing a buffer to an API that has a different error to check for insufficient buffer size vs [netstat](../netstat).

```
> networkparams.exe
Host Name: computer01
Domain Name:
DNS Servers:
        192.168.1.1
Node Type: Hybrid
Routing: disabled
ARP Proxy: disabled
DNS: disabled
```