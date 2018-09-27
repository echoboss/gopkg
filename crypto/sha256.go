package main

import (
	"crypto/sha256"
	"io"
	"fmt"
)

func main() {
	h := sha256.New()
	io.WriteString(h,"abc")
	fmt.Printf("%x",h.Sum(nil))
}
