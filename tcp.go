package main

// name: tcp.go
// date: 2017-01-17 17:30:52
// author: konghui@live.cn

type TCPHeader struct {
	SrcPort  uint16
	DstPort  uint16
	SeqNum   uint32
	AckNum   uint32
	Offset   uint8
	Flag     uint8
	Window   uint16
	Checksum uint16
}
