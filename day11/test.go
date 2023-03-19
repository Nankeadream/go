package main

import (
	"bytes"
	"day11/encryptAndDecode"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func main() {

	block, paddingText, err := encryptAndDecode.AES("6368616e676520746869732070617373", []byte("aaaaaaaaaaaaaaaa"))
	if err != nil {
		log.Error().Err(err).Send()
	}
	iv := encryptAndDecode.Iv([]byte("abcdef0123456789"))
	ciphertext := encryptAndDecode.Encrypt(block, paddingText, iv)
	fmt.Printf("密文：%x\n", ciphertext)
	decode := encryptAndDecode.Decode(block, ciphertext, iv)
	fmt.Printf("明文：%x， %[1]s\n", decode)
	//// 为了让大家理解算法，就没有封装成函数
	//password := "6368616e676520746869732070617373" // 密码,32个字符
	//plainText := []byte("exampleplaintext")        // 明文，16
	//fmt.Println(password, len(password))
	//// key密钥，自定义，可以是16、24
	//key, _ := hex.DecodeString(password)                                     // 按照16进制理解password fmt.Println(key, len(key))
	//fmt.Printf("明文长度=%d字节，密钥长度%d字节\n", len(plainText), len(key)) // 16 16
	//// 块必须为指定大小，不够就补齐
	//// 本次采用PKCS7Padding方式
	//blockSize := aes.BlockSize
	//fmt.Printf("默认分组大小为 %d字节\n", blockSize)
	//r := len(plainText) % blockSize // 余数
	//n := blockSize - r              // 待填充字节数
	//if n == blockSize {
	//	// 正好满足分组字节要求，追加一个块，每个字节填充块大小
	//	fmt.Printf("正好满足分组块大小要求，追加一个块(%d个字节)\n", n)
	//} else {
	//	fmt.Printf("不满足块大小要求，需要补充%d字节\n", n)
	//}
	//padding := bytes.Repeat([]byte{byte(n)}, n)
	//fmt.Printf("%d, %v\n", len(padding), padding)
	//paddingText := append(plainText, padding...) // 补完后去加密 fmt.Printf("%d, %v\n", len(paddingText), paddingText)
	//block, err := aes.NewCipher(key)
	//if err != nil {
	//	panic(err)
	//}
	//// 加密
	//// CBC模式需要提供一个与第一分组计算的初始向量iv, iv字节数为块大小，这里取16字节
	//iv := []byte("abcdef0123456789")
	//enMode := cipher.NewCBCEncrypter(block, iv)
	//cipherText := make([]byte, len(paddingText))
	//enMode.CryptBlocks(cipherText, paddingText)
	//fmt.Printf("密文:%x\n", cipherText)
	//// 解密，使用密文cipherText解密
	//deMode := cipher.NewCBCDecrypter(block, iv)
	//text := make([]byte, len(cipherText))
	//deMode.CryptBlocks(text, cipherText)
	//padding1 := text[len(text)-1] // text中最后一个字节一定是补充的字节数
	//fmt.Printf("明文:%x, %[1]s\n", text[:len(text)-int(padding1)])

	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	fmt.Printf("%T", w)
	fmt.Println(w)
	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	rwc = new(bytes.Buffer)
}
