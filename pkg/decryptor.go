package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"os"
)

func Decrypt(path string, key []byte) {
	infile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	originalPath := path[:len(path)-len(ENCRYPTED_FILE_EXTENSION)]
	fmt.Println(originalPath)
	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Deattached nonce and decrypt
	nonce := infile[:gcm.NonceSize()]
	infile = infile[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, infile, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	// Writing decryption content
	err = os.WriteFile(originalPath, plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
	os.Remove(path)
}
