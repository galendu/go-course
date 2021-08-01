package aes_test

import (
	"fmt"
	"testing"

	"gitee.com/infraboard/go-course/day9/aes"
)

func TestAES(t *testing.T) {
	plain := "The text need to be encrypt."
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"
	// 加密
	cipherByte, err := aes.Encrypt(plain, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %x\n", plain, cipherByte)
	// 解密
	plainText, err := aes.Decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x ==> %s\n", cipherByte, plainText)
}
