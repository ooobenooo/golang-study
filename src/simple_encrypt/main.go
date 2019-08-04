package main

import (
	"fmt"
	"encoding/base64"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	hello := "hello world"
	encodeByte := base64Encode([]byte(hello))
	fmt.Println(string(encodeByte))

	decodeByte, _ := base64Decode(encodeByte)
	fmt.Println(string(decodeByte))
}
