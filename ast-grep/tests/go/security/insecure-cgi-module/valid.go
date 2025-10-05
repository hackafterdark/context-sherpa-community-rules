package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func goodExample() {
	// This should not trigger the rule - no cgi usage
	block, err := des.NewCipher([]byte("sekritz"))
	if err != nil {
		panic(err)
	}
	plaintext := []byte("I CAN HAZ SEKRIT MSG PLZ")
	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[des.BlockSize:], plaintext)
	fmt.Println("Secret message is: %s", hex.EncodeToString(ciphertext))
}

func anotherGoodExample() {
	for _, arg := range os.Args {
		// This should not trigger the rule - no cgi usage
		fmt.Printf("%x - %s\n", md5.Sum([]byte(arg)), arg)
	}
}

func thirdGoodExample() {
	// This should not trigger the rule - no cgi usage
	cipher, err := rc4.NewCipher([]byte("sekritz"))
	if err != nil {
		panic(err)
	}
	plaintext := []byte("I CAN HAZ SEKRIT MSG PLZ")
	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)
	fmt.Println("Secret message is: %s", hex.EncodeToString(ciphertext))
}

func fourthGoodExample() {
	for _, arg := range os.Args {
		// This should not trigger the rule - no cgi usage
		fmt.Printf("%x - %s\n", sha1.Sum([]byte(arg)), arg)
	}
}
