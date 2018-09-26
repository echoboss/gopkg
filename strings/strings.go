package main

import (
	"fmt"
	"strings"
	"bytes"
)

func main() {
	fmt.Println(strings.Contains("Luyoo","oo"))
	fmt.Println(strings.Contains("LuYoo",""))
	fmt.Println(strings.ContainsAny("ContainsAny","a"))
	fmt.Println(strings.ContainsRune("你好",'你'))
	fmt.Println(strings.HasPrefix("http://a.com","http"))
	fmt.Println(strings.HasSuffix("a.jpg",".jpeg"))
	fmt.Println(strings.EqualFold("go","GO"))

	fmt.Println(strings.Count("Luyoo","o"))
	fmt.Println(strings.Count("Luyoo",""))
	fmt.Println(strings.Index("LuYoo","w"))
	fmt.Println(strings.IndexAny("Luyoo","ao"))

	fmt.Println(strings.IndexRune("Luyoo",'y'))
	index := func(r rune) bool {
		if r > 'b' {
			return true
		}
		return false
	}

	fmt.Println(strings.IndexFunc("abc",index))
	fmt.Println(strings.LastIndexFunc("abcda", index))
	fmt.Println(strings.Repeat("--",20))


	fmt.Println(strings.Join([]string{"hello","world"}," "))

	fmt.Printf("%q\n",strings.Split("Luyoo",""))
	fmt.Println(strings.SplitN("Luyoo","u",1))
	fmt.Println(strings.SplitAfter("a.jpg",".",))
	fmt.Println(strings.SplitAfterN("a.jpg",".",-1))
	fmt.Println(strings.Fields("I am Luyoo"))

	split := func(r rune) bool {
		if r > 'b' {
			return true
		}
		return false
	}

	fmt.Printf("%q\n",strings.FieldsFunc("abc",split))


	fmt.Println(strings.Title("luyoo"))
	fmt.Println(strings.ToLower("LuYoo"))
	fmt.Println(strings.ToTitle("lUyOO"))
	fmt.Println(strings.ToUpper("luyoo"))

	fmt.Println(strings.Trim("Luyoo","o"))
	fmt.Println(strings.TrimLeft("Luyoo","L"))
	fmt.Println(strings.TrimRight("Luyoo","o"))

	trimF := func(r rune) bool {
		if r < 'Z' {
			return true
		}
		return false
	}

	fmt.Println(strings.TrimFunc("LuYoo",trimF))
	fmt.Println(strings.TrimLeftFunc("LuYoo",trimF))
	fmt.Println(strings.TrimRightFunc("LuYoO",trimF))

	fmt.Printf("%q\n",strings.TrimSpace("\n\t\r abc\t\r"))

	mapRune := func(r rune) rune {
		return r + 1
	}

	fmt.Println(strings.Map(mapRune,"abc路"))

	r := bytes.NewBuffer(nil)


	rp := strings.NewReplacer("x", "X", "y", "Y")
	fmt.Println(rp.Replace("luyoo"))
	rp.WriteString(r,"Luyoo")
	fmt.Println("aaaa",r)

	fmt.Println(strings.Replace("Luyoo","y","Y",-1))

	a := "Luyoo路"
	rb := strings.NewReader(a)
	fmt.Println(rb.Len())

	b := make([]byte,2)

	n, err := rb.Read(b)
	fmt.Println(n,err)
	fmt.Printf("%s\n",b)

	n, err = rb.ReadAt(b,4)
	fmt.Println(n,err,string(b))

	oneb,err  := rb.ReadByte()
	fmt.Println(string(oneb),err)

	oner,s, err := rb.ReadRune()
	fmt.Println(string(oner),s,err)

	rb.Seek(0,0)
	oneb,err  = rb.ReadByte()
	fmt.Println(string(oneb),err)

	rb.UnreadByte()

	oneb,err  = rb.ReadByte()
	fmt.Println(string(oneb),err)



}
