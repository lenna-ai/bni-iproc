package decrypt

import (
	"crypto/aes"
	"encoding/hex"
)


func DecryptAES(encryptedHex string) (string, error) {
	// Encrypted string in hexadecimal format
	// encryptedHex := "1D4D4C3BE25EB5CDB19BE4C8EA06E1B9"
	// AES key used for encryption
	//key := "pr0m1sE2015"

	// Decode the encrypted string from hexadecimal to byte array
	encrypted, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "",err
	}

	//arr := []int{-45, -57, -75, -95, 72, 110, 91, -52, 9, -22, 100, -84, -3, -19, -25, 73}

	// Convert to constant []byte slice
	keyBytes := []byte{
		0xd3, 0xc7, 0xb5, 0xa1, 0x48, 0x6e, 0x5b, 0xcc,
		0x09, 0xea, 0x64, 0xac, 0xfd, 0xed, 0xe7, 0x49,
	}
	if len(keyBytes) != 16 {
		return "AES-128 key must be 16 bytes",err
	}

	// Create cipher block for AES
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "AES-128 key must be 16 bytes",err
	}

	// Decrypt the encrypted data
	decrypted := make([]byte, len(encrypted))
	block.Decrypt(decrypted, encrypted)

	// Trim padding from decrypted data (assuming PKCS#7 or similar padding)
	decrypted = pkcs7Unpad(decrypted)

	// Convert decrypted data to string
	decryptedString := string(decrypted)

	return decryptedString,nil
}


// pkcs7Unpad removes PKCS#7 padding from decrypted data
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}