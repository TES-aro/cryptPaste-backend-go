package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func aesEncrypt(plaintext string , key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("error at ciper.BlockSize")
		fmt.Println(err)
		return nil, err
	}
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddingText := append([]byte(plaintext), bytes.Repeat([]byte{byte(padding)},padding)...)

	cipherText := make([]byte, aes.BlockSize+len(paddingText))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader ,iv)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block,iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:],paddingText)
	fmt.Println("\nencrypted blocks:")
	fmt.Println(cipherText)
	return cipherText, err

}
