package main

import (
	"bytes"
	"bufio"
	"fmt"
)

const str = "cache and name things"

var bs = make([]byte, 22)

func main() {
	//s := strings.NewReader(str)
	//
	//bufr := bufio.NewReader(s)
	//
	//bufr.Read(bs)
	//fmt.Printf("%s\n",bs)
	//fmt.Println(bufr.Buffered())

	//bufr := bufio.NewReaderSize(s,10)
	//bufr.Read(bs)
	//fmt.Printf("%s\n",bs)

	//bufr.Peek(2)0
	//fmt.Printf("%s\n",peek)

	//cont := make([]byte,10)
	//
	//n, err := bufr.Read(bs)
	//fmt.Println(n,err)
	//fmt.Println(string(bs))
	//n, err = bufr.Read(bs)
	//fmt.Println(n,err)

	//s.Reset("路遥")
	//bufr := bufio.NewReader(s)

	//body, err := bufr.ReadBytes('_')
	//fmt.Println(string(body),err)

	//r,n, err := bufr.ReadRune()
	//fmt.Println(string(r),n,err)

	////
	//wb := bytes.NewBuffer(nil)
	//w := bufio.NewWriter(wb)
	//w.Write([]byte("Fuck"))
	//w.WriteRune('哈')
	//
	//fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))
	//w.Flush()
	//fmt.Printf("%d:%s\n", len(wb.Bytes()), string(wb.Bytes()))

	r := bytes.NewBuffer([]byte("Abc"))
	w := bytes.NewBuffer(nil)

	rb := bufio.NewReader(r)
	wb := bufio.NewWriter(w)

	rw := bufio.NewReadWriter(rb, wb)

	rw.Write([]byte("Fuck"))
	rw.Flush()

	n, err := rw.Read(bs)
	fmt.Printf("%d,%v\n",n,err)
	fmt.Printf("%s\n",bs)

	fmt.Printf("%s",w)



}
