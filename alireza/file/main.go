package main

import (
	"fmt"
	"os"
)

// ===== Membuat File Baru ====
var path = "/home/alireza/.gvm/pkgsets/go1.14.3/go_alireza/go_novalagung/github.com/alireza/file/text.txt"

/* var path = "/home/alireza/.gvm/pkgsets/go1.14.3/go_alireza/go_novalagung/github.com/alireza/file/text.json"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada

	_, err := os.Stat(path)

	//buat file baru jika belum ada

	if os.IsNotExist(err) {
		file, err := os.Create(path)

		if isError(err) {
			return
		}

		defer file.Close()
	}

	fmt.Println("==== file berhasil dibuat")

} */

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// ===== Edit File ====

/* func writeFile() {
	// buka file dengan akses READ & WRITE

	file, err := os.OpenFile(path, os.O_RDWR, 0644)

	if isError(err) {
		return
	}

	defer file.Close()

	// tulis data ke File

	_, err = file.WriteString("halo\n")

	if isError(err) {
		return
	}

	_, err = file.WriteString("mari belajar Golang \n")

	if isError(err) {
		return
	}

	/// simpan perubahan

	err = file.Sync()

} */

/* func readFile() {
	// buka file

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)

	if isError(err) {
		return
	}

	defer file.Close()

	//baca File

	text := make([]byte, 1024)

	for {
		n, err := file.Read(text)

		if err != io.EOF {
			if isError(err) {
				break
			}
		}

		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil dibaca")

	fmt.Println(string(text))

} */

// ===== Menghapus File  =====
func deleteFile() {

	err := os.Remove(path)

	if isError(err) {
		return
	}

	fmt.Println("===> file berhasil di hapus")

}

func main() {
	fmt.Println("Edit file")
	// readfile
	deleteFile()
	//readFile()
	//writeFile()
	//createFile()

}
