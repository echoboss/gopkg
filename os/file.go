package main

import (
	"os"
	"fmt"
)

var cont = `
bcd
f hij
`

func main() {
	//err := os.Mkdir("demo",0755)
	//fmt.Println(err)
	//err := os.MkdirAll("demo/app",0755)
	//fmt.Println(err)

	//f,err := os.Create("os.go")
	//fmt.Println(err)
	//defer f.Close()
	//fmt.Println(f.Name(),f.Fd())
	//
	//n, err := f.Write([]byte(cont))
	//fmt.Println(n,err)
	//n, err = f.WriteAt([]byte("\nFuck"),9)
	//fmt.Println(n,err)
	//n,err = f.WriteString("\nabc")
	//fmt.Println(n,err)
	//f.Sync()
	//
	//nf := os.NewFile(f.Fd(),"abc.txt")
	//nf.Sync()
	//defer nf.Close()

	//f, err := os.Open("os.go")
	//fmt.Println(err)
	f, err := os.OpenFile("os.go",os.O_CREATE|os.O_RDWR,0644)
	fmt.Println(err)

	//b := make([]byte,3)
	//f.Read(b)
	//fmt.Printf("%q\n",b)
	defer f.Close()

	//n, _ := f.ReadAt(b,6)
	//fmt.Printf("%q\n",b[:n])


	//f.Seek(1,1)
	//n, _ = f.Read(b)
	//fmt.Printf("%q\n",b[:n])

	//f.Write([]byte("Luyoo"))

	//f.Truncate(2)
	//
	//f.Read(b)
	//fmt.Printf("%s",b)

	//info, err := f.Stat()
	//fmt.Println(err)
	//fmt.Println(info.Name())
	//fmt.Println(info.Size())
	//fmt.Println(info.IsDir())
	//fmt.Println(info.ModTime())
	//fmt.Println(info.Mode().IsRegular())
	//


	//_, err = os.Open("a.xxx")
	//fmt.Println(os.IsExist(err))
	//fmt.Println(os.IsNotExist(err))

	//_, err = os.Open("os.go")
	//fmt.Println(os.IsPermission(err))
	//fmt.Println(os.IsPathSeparator('/'))
	os.Remove("os.go")









}