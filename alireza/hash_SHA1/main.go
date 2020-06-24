package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func doHashUsingSalt(text string) (string, string) {

	salt := fmt.Sprintf("%d", time.Now().UnixNano())

	saltedText := fmt.Sprintf("text:'%s', salt: %s ", text, salt)

	fmt.Println(saltedText)

	sha := sha1.New()

	sha.Write([]byte(saltedText))

	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {

	text := "secret"

	fmt.Printf("original: %s\n\n", text)

	hashed1, salt1 := doHashUsingSalt(text)
	fmt.Printf("hashed1: %s\n\n", hashed1)

	hashed2, salt2 := doHashUsingSalt(text)
	fmt.Printf("hashed2: %s\n\n", hashed2)

	hashed3, salt3 := doHashUsingSalt(text)
	fmt.Printf("hashed3: %s\n\n", hashed3)

	_, _, _ = salt1, salt2, salt3

	/* text := "this is secret"

	sha := sha1.New()

	sha.Write([]byte(text))

	encrypted := sha.Sum(nil)

	encryptedString := fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString) */

	// ==== Metode Salt pada Hash SHA1 ====

}
