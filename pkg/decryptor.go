package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func Decrypt(path string) {
	infile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	//defer infile.Close()

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

	originalPath := path[:len(path)-4]
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
	//fi, err := infile.Stat()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//iv := make([]byte, block.BlockSize())
	//msgLen := fi.Size() - int64(len(iv))
	//_, err = infile.ReadAt(iv, msgLen)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//outfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer outfile.Close()
	//
	//// The buffer size must be multiple of 16 bytes
	//buf := make([]byte, 1024)
	//stream := cipher.NewCTR(block, iv)
	//for {
	//	n, err := infile.Read(buf)
	//	if n > 0 {
	//		// The last bytes are the IV, don't belong the original message
	//		if n > int(msgLen) {
	//			n = int(msgLen)
	//		}
	//		msgLen -= int64(n)
	//		stream.XORKeyStream(buf, buf[:n])
	//		// Write into file
	//		outfile.Write(buf[:n])
	//	}
	//
	//	if err == io.EOF {
	//		break
	//	}
	//
	//	if err != nil {
	//		log.Printf("Read %d bytes: %v", n, err)
	//		break
	//	}
	//}
}
