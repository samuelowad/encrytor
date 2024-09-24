package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func Encrypt(path string, key []byte) {
	infile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce err: %v", err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, infile, nil)

	// Writing encryption content
	err = os.WriteFile(path+ENCRYPTED_FILE_EXTENSION, ciphertext, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
	os.Remove(path)
}
