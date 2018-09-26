package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Printf("%q\n",strconv.Itoa(233))
	fmt.Printf("%q\n",strconv.FormatBool(true))
	fmt.Printf("%q\n",strconv.FormatInt(10,16))
	fmt.Printf("%q\n",strconv.FormatFloat(3.14,'f',2,64))


	fmt.Println(strconv.Atoi("1"))
	fmt.Println(strconv.ParseInt("1000",2,10))
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseFloat("3.14",64))

	b := []byte("a")
	fmt.Printf("%q\n",strconv.AppendBool(b,true))
	fmt.Printf("%q\n",strconv.AppendFloat(b,3.14,'e',2,64))

	fmt.Printf("%q\n",strconv.AppendInt(b,233,10))
	fmt.Printf("%q\n",strconv.AppendUint(b,233,36))
	fmt.Printf("%q\n",strconv.AppendQuote(b,"fuck"))

	fmt.Printf("%q\n",strconv.AppendQuoteRune(b,'你'))


	fmt.Println(strconv.Quote("aaaa"),"aaaa")

	fmt.Println(strconv.Unquote("\"aaaaa\""))

	fmt.Println(strconv.CanBackquote("C:\\Win32")) // true
	fmt.Println(strconv.CanBackquote("C:\\Win32/add")) // false
	fmt.Println(strconv.CanBackquote("C:\\Win32`"))  // false

	fmt.Println(strconv.IsPrint(' '), strconv.IsPrint('\n')) // true false
}
