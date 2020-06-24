package main

import (
	"flag"
	"fmt"
)

func main() {
	// ==== arguments =====

	/*  argsRaw := os.Args

	    fmt.Printf("-> %#v\n", argsRaw)

	    args := argsRaw[1:]
	    fmt.Printf("-> %#v\n", args) */

	// ==== Flags =====

	/* name := flag.String("name", "anonymous", "your name")

	age := flag.Int64("age", 24, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)

	fmt.Printf("age\n: %d\n", age) */

	// ==== Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data ====

	data1 := flag.String("name", "anonymous", "type your name")

	fmt.Println(*data1)
}
