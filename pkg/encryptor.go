package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func Encrypt(file string) {
	fileText, err := os.ReadFile(file)
	if err != nil {
		//return err
		log.Fatal(err)
	}

	// The key should be 16 bytes (AES-128), 24 bytes (AES-192) or
	// 32 bytes (AES-256)
	key, err := os.ReadFile("./shared/key/key.txt")
	if err != nil {
		log.Fatal(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, fileText, nil)
	// Save back to file
	err = os.WriteFile(file+".bin", ciphertext, 0777)
	if err != nil {
		//return err
		log.Panic(err)
	}
	os.Remove(file)

	//return nil
}
