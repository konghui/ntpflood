package main

// name: ip.go
// date: 2017-01-17 17:27:40
// author: konghui@live.cn

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var IPV4 = 4
var IPV4_HDR_LEN = 5

type IPHeader struct {
	version   uint8
	tos       uint8
	length    uint16
	id        uint16
	offset    uint16
	ttl       uint8
	protoType uint8
	checkSum  uint16
	srcAddr   uint32
	dstAddr   uint32
}

func NewIPHeader(protoType uint8, srcAddr, dstAddr string) (header *IPHeader, err error) {
	header = &IPHeader{version: uint8((IPV4 << 4) + IPV4_HDR_LEN), tos: 0, length: 0, id: 0, offset: 0, ttl: 255, protoType: protoType, checkSum: 0}
	header.srcAddr, err = inet_addr(srcAddr)
	if err != nil {
		return
	}
	header.dstAddr, err = inet_addr(dstAddr)
	if err != nil {
		return
	}
	return
}

// (this *IPHeader) ...
func (this *IPHeader) String() (s string) {
	s = fmt.Sprintf("version:%d, hdr_len:%d, tos:%d, length:%d, id: %d, offset:%d, ttl:%d, protoType:%d, checksum: %d, src addr:%s, dst addr:%s\n", (this.version>>4)&0xF, this.version&0xF, this.tos, this.length, this.id, this.offset, this.ttl, this.protoType, this.checkSum, inet_ntoa(this.srcAddr), inet_ntoa(this.dstAddr))
	return
}

// NetworkByte
// return the object with network byte(BigEndian)
func (this *IPHeader) NetworkByte() (b []byte, err error) {
	b, err = GetBigEndianData(this)
	return
}

func (this *IPHeader) SetCheckSum(data []byte) {

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, this)
	binary.Write(&buffer, binary.BigEndian, data)
	this.checkSum = checkSum(buffer.Bytes())
}
