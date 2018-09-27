package main

import (
	"crypto/md5"
	"io"
	"fmt"
)

func main() {
	h := md5.New()

	io.WriteString(h,"https://www.jianshu.com/u/e372dabaafaf")
	fmt.Printf("%x\n",h.Sum(nil))
}
