package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// name: common.go
// date: 2017-01-17 17:25:04
// author: konghui@live.cn

func GetBigEndianData(data interface{}) (b []byte, err error) {
	var (
		buffer bytes.Buffer
	)
	err = binary.Write(&buffer, binary.BigEndian, data)
	b = buffer.Bytes()
	return
}

func inet_ntoa(addr uint32) (ip string) {
	var ipList = make([]uint8, 4)

	ipList[3] = uint8(addr & 0xFF)
	ipList[2] = uint8((addr >> 8) & 0xFF)
	ipList[1] = uint8((addr >> 16) & 0xFF)
	ipList[0] = uint8((addr >> 24) & 0xFF)
	ip = fmt.Sprintf("%d.%d.%d.%d", ipList[0], ipList[1], ipList[2], ipList[3])
	return
}

func inet_addr(ipaddr string) (ret uint32, err error) {

	var ip uint64
	var ipList = make([]uint64, 4)

	for i, ipStr := range strings.Split(ipaddr, ".") {
		ip, err = strconv.ParseUint(ipStr, 10, 64)
		ipList[i] = ip
	}
	ret = uint32(ipList[0]<<24 + ipList[1]<<16 + ipList[2]<<8 + ipList[3])

	return
}

func htons(port uint16) (p uint16) {
	p = port<<8 + port>>8
	return
}

func checkSum(data []byte) uint16 {
	fmt.Println(data)
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
