package decrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

func DecryptAesFrontend(encryptText string ) (string, error) {
	secret_key_login := os.Getenv("SECRET_KEY_LOGIN")
	key := []byte(secret_key_login)

	// result, err := AesEncrypt([]byte("bni1234/"), key)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(base64.StdEncoding.EncodeToString(result)) //VRMzYJwT5xx0bvqud3Np+g==

	// r, _ := base64.StdEncoding.DecodeString("TIIKG6mAfA0gNt0CSuFccA==") // use CryptoJs encrypted
	r, _ := base64.StdEncoding.DecodeString(encryptText) // use CryptoJs encrypted
	//r := result  // decrypt go encrypted
	origData, err := AesDecrypt(r, key)
	if err != nil {
		return "", err
	}
	fmt.Println(string(origData))
	return string(origData), nil
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	iv := []byte(os.Getenv("SECRET_KEY_LOGIN_IV"))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	secret_key_login_iv := os.Getenv("SECRET_KEY_LOGIN_IV")
	iv := []byte(secret_key_login_iv)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
