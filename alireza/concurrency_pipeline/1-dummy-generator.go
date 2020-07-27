/* package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}

	return string(b)

}

func generatedFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(contentLength)
		err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "file created")
		}
	}
	log.Printf("%d of total files created ", totalFile)
}

func main() {
	log.Print("Start")

	start := time.Now()

	generatedFiles()

	duration := time.Since(start)

	log.Println("done in", duration.Seconds(), "seconds")

}
 */