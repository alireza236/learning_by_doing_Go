package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	/* data := "John wick"

	encodeString := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("encoded:", encodeString)

	decodeByte, _ := base64.StdEncoding.DecodeString(encodeString)

	decodeString := string(decodeByte)

	fmt.Println("decoded:", decodeString) */

	// ===== Penerapan fungsi Decoded dan Encoded ====

	/* data := "Ali Reza"

	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))

	base64.StdEncoding.Encode(encoded, []byte(data))

	encodedString := string(encoded)

	fmt.Println(encodedString)

	decoded := make([]byte, base64.StdEncoding.EncodedLen(len(encoded)))

	_, err := base64.StdEncoding.Decode(decoded, encoded)

	if err != nil {
		fmt.Println(err.Error())
	}

	decodedString := string(decoded)

	fmt.Println(decodedString) */

	// ===== Encode dn Decode pada URl ====

	data := "https://kalipare.com"

	encodedString := base64.URLEncoding.EncodeToString([]byte(data))

	fmt.Println(encodedString)

	decodedByte, _ := base64.URLEncoding.DecodeString(encodedString)

	decodedString := string(decodedByte)

	fmt.Println(decodedString)

}
