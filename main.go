package main

import (
	"fmt"
	"github/Ghun2/learn-go-decorator/cipher"
	"github/Ghun2/learn-go-decorator/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

type SendComponent struct {}

func (self *SendComponent) Operator(data string) {
	// Send Data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct {}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	sender := &EncryptComponent{
		key: "asdf",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("HI HELLO")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "asdf",
			com: &ReadComponent{},
		},
	}
	receiver.Operator(sentData)
	fmt.Println(recvData)
}


