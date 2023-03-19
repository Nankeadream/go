package encryptAndDecode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

//函数封装

func AES(password string, plainText []byte) (cipher.Block, []byte, error) {
	fmt.Println(password, len(password))
	//key密钥，自定义，可以是16，24，32字节
	key, _ := hex.DecodeString(password)
	fmt.Println(key, len(key))
	fmt.Printf("明文长度=%d字节,密钥长度%d字节\n", len(plainText), len(key))
	//块必须为指定大小，不够就不起
	//本次采用PKCS7Padding方法
	blockSize := aes.BlockSize
	fmt.Printf("默认分组大小为%d字节\n", blockSize)
	r := len(plainText) % blockSize //余数
	n := blockSize - r              //待填充字节书
	if n == blockSize {
		//正好满足分组字节要求，追加一个块，每个字节填充块大小
		fmt.Printf("正好满足分组块大小要求，追加一个块(%d个字节)\n", n)
	} else {
		fmt.Printf("不满足块大小要求，需要补充%d字节\n", n)
	}
	padding := bytes.Repeat([]byte{byte(n)}, n)
	fmt.Printf("%d,%v\n", len(padding), padding)
	paddingText := append(plainText, padding...) //补完后去加密
	fmt.Printf("%d, %v\n", len(paddingText), paddingText)
	block, err := aes.NewCipher(key)
	return block, paddingText, err
}

// Iv 设置IV值
func Iv(iv []byte) []byte {
	return iv
}

// Encrypt 加密
// CBC模式需要提供一个于第一分组计算的初始向量iv，iv字节数为块大小，这里取16字节
func Encrypt(block cipher.Block, paddingText []byte, iv []byte) []byte {
	enMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddingText))
	enMode.CryptBlocks(cipherText, paddingText)
	return cipherText
}

// Decode 解密，使用密文ciphertext解密
func Decode(block cipher.Block, cipherText []byte, iv []byte) []byte {
	deMode := cipher.NewCBCDecrypter(block, iv)
	text := make([]byte, len(cipherText))
	deMode.CryptBlocks(text, cipherText)
	padding := text[len(text)-1] //text中最后一个字节一定是补充的字节数
	return text[:len(text)-int(padding)]
}
