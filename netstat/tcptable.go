package main

import (
	"encoding/binary"
	"net"
	"unsafe"
)

type TableRow struct {
	Local  *net.TCPAddr
	Remote *net.TCPAddr
	State  TcpState
	PID    int
}

func GetTCPTable() ([]TableRow, error) {
	var tableSize uint32
	var tcpTable []byte
	// query for the buffer size by sending null/0
	ret := _GetExtendedTcpTable(nil, &tableSize, true, _AF_INET4, _TCP_TABLE_OWNER_PID_ALL, 0)
	for {
		if ret == _NO_ERROR {
			break
		}
		if ret == _ERROR_INSUFFICIENT_BUFFER {
			tcpTable = make([]byte, int(tableSize))
			ret = _GetExtendedTcpTable(&tcpTable[0], &tableSize, true, _AF_INET4, _TCP_TABLE_OWNER_PID_ALL, 0)
			continue
		}
		return nil, ret
	}
	return convertToTableRows(tcpTable), nil
}

func convertToTableRows(raw []byte) (result []TableRow) {
	if len(raw) == 0 {
		return
	}
	table := (*_MIB_TCPTABLE_OWNER_PID)(unsafe.Pointer(&raw[0]))
	n := int(table.NumEntries)
	if n == 0 {
		return
	}
	result = make([]TableRow, n)
	rows := unsafe.Slice(&table.Table[0], table.NumEntries)
	for i, row := range rows {
		result[i] = TableRow{
			Local:  convertToTCPv4Addr(row.LocalAddr, row.LocalPort),
			Remote: convertToTCPv4Addr(row.RemoteAddr, row.RemotePort),
			PID:    int(row.OwningPID),
			State:  TcpState(row.State),
		}
	}
	return
}

func convertToTCPv4Addr(i uint32, p uint32) *net.TCPAddr {
	var ip net.IP
	if i > 0 { // convert from network byte order
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, i)
		ip = bytes
	}
	var port int
	if p > 0 { // convert from network byte order
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, uint16(p))
		port = int(binary.BigEndian.Uint16(bytes))
	}
	return &net.TCPAddr{
		IP:   ip,
		Port: port,
	}
}
