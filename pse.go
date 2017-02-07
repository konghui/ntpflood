package main

// name: pse.go
// date: 2017-01-17 17:28:59
// author: konghui@live.cn

import (
	"fmt"
)

type PseHeader struct {
	srcAddr   uint32
	dstAddr   uint32
	zero      uint8
	protoType uint8
	length    uint16
}

func NewPseHeader(srcAddr, dstAddr string, protoType uint8) (p *PseHeader, err error) {
	p = &PseHeader{zero: 0, protoType: protoType, length: 0}
	p.srcAddr, err = inet_addr(srcAddr)
	if err != nil {
		return
	}
	p.dstAddr, err = inet_addr(dstAddr)
	if err != nil {
		return
	}
	return
}

func (this *PseHeader) String() (s string) {
	s = fmt.Sprintf("src addr:%s, dst addr:%s, zero:%d, protoType:%d, length:%d\n", inet_ntoa(this.srcAddr), inet_ntoa(this.dstAddr), this.zero, this.protoType, this.length)
	return
}

func (this *PseHeader) NetworkByte() (b []byte, err error) {
	b, err = GetBigEndianData(this)
	return
}
