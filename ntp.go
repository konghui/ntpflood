package main

// name: ntp.go
// date: 2017-01-17 13:42:10
// author: konghui@live.cn

import "fmt"

var VERSION = 2

type NTPHeader struct {
	version         uint8
	stratum         uint8
	poll            uint8
	precision       uint8
	syncDistance    uint32
	syncDispersion  uint32
	referenceId     uint32
	referTimestamp  uint32
	originTimestamp uint32
	recvTimestamp   uint32
	transTimestamp  uint32
}

type NTPResponseHeader struct {
	flag      uint8
	auth      uint8
	poll      uint8
	precision uint8
	num       uint16
	length    uint16
}

func NewNTPHeader(li uint8, version uint8, mode uint8, stratum uint8, poll uint8, precision uint8) (n *NTPHeader) {
	n = &NTPHeader{version: li<<6 + version<<3 + mode, stratum: stratum, poll: poll, precision: 42}
	return
}

func (this *NTPHeader) NetworkByte() (b []byte, err error) {
	b, err = GetBigEndianData(*this)
	return
}

func NewMonListRequest() (b []byte, err error) {
	var ntp = NewNTPHeader(0, 2, 7, 0, 3, 42)
	b, err = ntp.NetworkByte()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return
}

/*func SendMonListRequest() {
	var count = 0
	var buffer = make([]byte, 512)
	var ntp = NewNTPHeader(0, 2, 7, 0, 3, 42)
	b, err = ntp.NetworkByte()
	if err != nil {
		fmt.Printf(err.Error())
	}
	conn, err := net.Dial("udp", "223.82.209.82:123")
	if err != nil {
		fmt.Printf(err.Error())
	}
	_, err = conn.Write(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		fmt.Printf("read\n")
		count, err = conn.Read(buffer)
		fmt.Print(cont)
		if cont == 0 {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(buffer)
	}
}*/
