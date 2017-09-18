package main

import (
	"bytes"
	"os"
	"encoding/json"
	"encoding/binary"
)

type Message struct {
	Player string
	Type string
	Url string
	Dest string
	Name string
}

type Response struct {
	Error bool
	Message string
}
func Send(err bool, msg string) {
	byteMsg := encodeMessage(err, msg)
	var msgBuf bytes.Buffer
	writeMessageLength(byteMsg)
	msgBuf.Write(byteMsg)
	msgBuf.WriteTo(os.Stdout)
}

func DecodeMessage(msg []byte) Message {
	var aMessage Message
	json.Unmarshal(msg, &aMessage)
	return aMessage
}

func encodeMessage(err bool, msg string) []byte {
	message := Response {
		Error: err,
		Message: msg,
	}
	return dataToBytes(message)
}

func dataToBytes(msg Response) []byte {
	byteMsg, _ := json.Marshal(msg)
	return byteMsg
}

func writeMessageLength(msg []byte) {
	binary.Write(os.Stdout, binary.LittleEndian, uint32(len(msg)))
}

func readMessageLength(msg []byte) int {
	var length uint32
	buf := bytes.NewBuffer(msg)
	binary.Read(buf, binary.LittleEndian, &length)
	return int(length)
}
