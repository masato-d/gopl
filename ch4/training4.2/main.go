package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	hashOpt = flag.String("h", "sha256", "hash type option")
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()

	flag.Parse()

	switch (*hashOpt) {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
	default:
		fmt.Println("invalid hash option")
	}
}