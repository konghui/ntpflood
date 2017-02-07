package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

type UDP struct {
	udpHeader *UDPHeader
	psdHeader *PseHeader
	ipHeader  *IPHeader
	fd        int
}

func NewUDP(srcAddr, dstAddr string, sport, dport uint16) (udp *UDP, err error) {
	udp = new(UDP)
	udp.psdHeader, err = NewPseHeader(srcAddr, dstAddr, syscall.IPPROTO_UDP)
	if err != nil {
		return
	}
	udp.udpHeader = NewUDPHeader(sport, dport)
	udp.ipHeader, err = NewIPHeader(syscall.IPPROTO_UDP, srcAddr, dstAddr)
	if err != nil {
		return
	}
	udp.fd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_UDP)
	if err != nil {
		return
	}
	err = syscall.SetsockoptInt(udp.fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		return
	}
	return
}

// (this *UDP)CheckSum ...
func (this *UDP) checkSum(data []byte) (err error) {
	var (
		buffer bytes.Buffer
		b      []byte
	)

	this.udpHeader.checkSum = 0 // must set checksum = 0 before caculate

	this.udpHeader.length = uint16(unsafe.Sizeof(*this.udpHeader)) + uint16(len(data)) // set the length of udp header
	this.psdHeader.length = uint16(unsafe.Sizeof(*this.udpHeader)) + uint16(len(data))

	b, err = this.psdHeader.NetworkByte()
	fmt.Printf("psd:%x\n", b)
	if err != nil {
		return
	}
	buffer.Write(b)
	b, err = this.udpHeader.NetworkByte()
	if err != nil {
		return
	}
	fmt.Printf("udp:%x\n", b)
	buffer.Write(b)
	buffer.Write(data)
	fmt.Printf("pse:%s, udp:%s\n", this.psdHeader, this.udpHeader)
	fmt.Printf("buffer:%x\n", buffer.Bytes())
	this.udpHeader.checkSum = checkSum(buffer.Bytes())
	return
}

func (this *UDP) Send(data []byte) (err error) {
	var (
		addr   syscall.SockaddrInet4
		ip     string
		b      []byte
		n      uint64
		buffer bytes.Buffer
	)
	ip = inet_ntoa(this.ipHeader.dstAddr)
	for i, v := range strings.Split(ip, ".") {
		n, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			return
		}

		addr.Addr[i] = uint8(n)
	}

	addr.Port = int(this.udpHeader.dstPort)

	err = this.checkSum(data)
	if err != nil {
		return
	}

	this.ipHeader.SetCheckSum(b)
	b, err = this.ipHeader.NetworkByte()
	if err != nil {
		return
	}
	buffer.Write(b)
	b, err = this.udpHeader.NetworkByte()
	if err != nil {
		return
	}

	buffer.Write(b)
	buffer.Write(data)
	fmt.Println(this.udpHeader)
	err = syscall.Sendto(this.fd, buffer.Bytes(), 0, &addr)
	if err != nil {
		return
	}
	return
}

// (this *UDP)SendAll ...
func (this *UDP) SendString(s string) (err error) {
	err = this.Send([]byte(s))
	return
}

func (this *UDP) Close() {
	syscall.Shutdown(this.fd, syscall.SHUT_RDWR)
}

type UDPHeader struct {
	srcPort  uint16
	dstPort  uint16
	length   uint16
	checkSum uint16
}

func NewUDPHeader(sport, dport uint16) (udp *UDPHeader) {
	udp = &UDPHeader{srcPort: sport, dstPort: dport}

	return
}

func (this *UDPHeader) String() (s string) {
	s = fmt.Sprintf("src port:%d dst port:%d length:%d checksum:%x", this.srcPort, this.dstPort, this.length, this.checkSum)
	return
}

func (this *UDPHeader) NetworkByte() (b []byte, err error) {
	b, err = GetBigEndianData(*this)
	return

}

func SendTo(src, dst string, sport, dport uint16, data []byte) (err error) {
	var udp *UDP
	udp, err = NewUDP(src, dst, sport, dport)
	if err != nil {
		return
	}

	err = udp.Send(data)
	if err != nil {
		return
	}
	udp.Close()
	return
}
