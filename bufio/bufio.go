package main

import (
	"strings"
	"fmt"
	"bufio"
)

const str = "cache and name things"

var bs = make([]byte,21)

func main() {
	s := strings.NewReader(str)
	//
	bufr := bufio.NewReader(s)
	//
	//bufr.Read(bs)
	//fmt.Printf("%s\n",bs)
	//fmt.Println(bufr.Buffered())

	//bufr := bufio.NewReaderSize(s,10)
	//bufr.Read(bs)
	//fmt.Printf("%s\n",bs)

	bufr.Peek(2)
	//fmt.Printf("%s\n",peek)

	//cont := make([]byte,10)
	//
	bufr.Read(bs)
	fmt.Println(string(bs))


}
