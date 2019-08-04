package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var iv = []byte("1234567890123456") //16个字符长度

func main() {
	plaintext := []byte("hello world")
	
	key_text := "abcdefghijklmnopqrstuvwxyz123456" // 至少32位

	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Println("aes error:", err)
		os.Exit(1)
	}

	cfb := cipher.NewCFBEncrypter(c, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%x\n", ciphertext)

	cfb_decrypt := cipher.NewCFBDecrypter(c, iv)
	plaintext_copy := make([]byte, len(plaintext))
	cfb_decrypt.XORKeyStream(plaintext_copy, ciphertext)
	fmt.Println(string(plaintext_copy))
}
