package main

import (
	"strings"
	"fmt"
	"io"
	"os"
	"io/ioutil"
)


var bs = make([]byte, 20)

func main() {

	r := strings.NewReader("MS微软")

	fmt.Println(io.CopyN(os.Stdout, r, 2))
	fmt.Println(io.Copy(os.Stdout, r))

	r.Reset("Amazon亚马逊")

	//if n, err := io.ReadAtLeast(r,bs,20); err == io.ErrUnexpectedEOF {
	//	fmt.Printf("read %d want 20\n",n)
	//	bs = bs[:n]
	//	fmt.Println(string(bs))
	//}

	//r.Reset("Google谷歌")
	//fmt.Println(io.ReadFull(r,bs))

	//r.Reset("Facebook脸书")
	//
	//s := io.NewSectionReader(r, 2, 20)
	//
	//n, err := s.Read(bs)
	//if err == io.EOF {
	//	fmt.Println(n,err)
	//	fmt.Printf("%s", bs)
	//
	//
	//}


	//r.Reset("gopher")
	//
	//tee := io.TeeReader(r,os.Stdout)
	//fmt.Println(tee.Read(bs))


	//fmt.Println(io.WriteString(os.Stdout,"abc"))

	//f, _ := os.Open("io.go")
	//lr := io.LimitedReader{R:f,N:20}
	//lr.Read(bs)
	//fmt.Println(lr.N,string(bs))

	//r.Reset("Apple苹果")
	//m := io.MultiReader(r,f)
	//var cont []byte
	//
	//for {
	//	n, err := m.Read(bs)
	//	if err == io.EOF {
	//		break
	//	}
	//	cont = append(cont,bs[:n]...)
	//}
	//fmt.Printf("%s",cont)

	//mw := io.MultiWriter(os.Stdout,f)
	//mw.Write(bs)
	//fmt.Println(bs)


	//pr,pw := io.Pipe()
	//go func() {
	//	defer pw.Close()
	//	pw.Write([]byte("Hello Abc"))
	//}()
	//
	//pr.Read(bs)
	//fmt.Printf("%s",bs)


	//f, _ := os.Open("io.go")
	//body, _ := ioutil.ReadAll(f)
	//fmt.Printf("%s",body)

	//
	//files, _ := ioutil.ReadDir("/tmp")
	//for _, f := range files {
	//	fmt.Println(f.Name())
	//}
	//
	//body, _ := ioutil.ReadFile("io.go")
	//fmt.Printf("%s",body)


	//n, _ := ioutil.TempDir(".","temp_")
	//
	//fmt.Println(n)


	n, _ := ioutil.TempFile(".","temp_")
	defer n.Close()
	//n.Write([]byte("abc"))
	n.WriteString("hello")
	fmt.Println(n.Name())






}
