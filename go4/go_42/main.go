package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	bitPtr := flag.Int("bits", 256, "256, 384, 512 bit")

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	flag.Parse()
	switch *bitPtr {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(input))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(input))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(input))
	default:
		fmt.Fprintln(os.Stderr, "Usage: Go_Practice_4_2 [-bits=256|384|512]")
	}
}
