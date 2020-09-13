package main

/*
IPHLPAPI_DLL_LINKAGE DWORD GetNetworkParams(
  PFIXED_INFO pFixedInfo,
  PULONG      pOutBufLen
);
*/

//sys _GetNetworkParams(fixedInfo *byte, outBufLen *uint32) (ret syscall.Errno)  = iphlpapi.GetNetworkParams
