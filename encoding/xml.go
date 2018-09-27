package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Loc struct {
	Country  string
	Province string
}

type City struct {
	Loc
	ID     int    `xml:"-"`                // 不输出
	People int    `xml:"people,omitempty"` // 值为空时不输出
	Name   string `xml:"name"`
}

func main() {

	c := City{ID: 1, Name: "西安"}
	c.Loc = Loc{"中国", "陕西"}
	b, _ := xml.Marshal(c)
	b, _ = xml.MarshalIndent(c, "", "    ")
	fmt.Println(string(b))
	s := `<x>233</x>`
	xml.Escape(os.Stdout, []byte(s))
}
