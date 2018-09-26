package main

import (
	"os"
	"fmt"
)

func main() {
	//oldpwd, _ := os.Getwd()
	//os.Chdir("/tmp")
	//
	//newpwd, _ := os.Getwd()
	//
	//fmt.Println(oldpwd,newpwd)
	//
	//file := path.Join(oldpwd,"os.go")
	//f, _ := os.Stat(file)
	//
	//os.Chmod(file,0664)
	//
	//of, _ := os.Stat(file)
	//fmt.Println(f.Mode(),of.Mode()&0777)
	//os.Chown(file,501,501)
	//
	//os.Chtimes(file,time.Now(),time.Now())
	s := "$sh is better than $$go"                  // $$ 不再解析
	fmt.Println(os.Expand(s, mapping))

	//os.Setenv("port","50001")
	//port := os.Getenv("port")
	//fmt.Println(port)
	//fmt.Println(os.ExpandEnv("go path is : $GOPATH"))
	//fmt.Println(os.ExpandEnv("go root path is: $$GOROOT"))
	//fmt.Println(os.Environ())
	//os.Clearenv()

	//fmt.Println(os.Getuid())
	//fmt.Println(os.Getgid())
	//
	//fmt.Println(os.Getgroups())

	//fmt.Println("go goroutine: ",os.Getpid())
	//
	//go func() {
	//	fmt.Println("parent goroutine:", os.Getppid())
	//}()
	//
	//time.Sleep(100 * time.Millisecond)

	fmt.Println(os.Getpagesize())
	fmt.Println(os.Hostname())
	fmt.Println(os.TempDir())


}


func mapping(s string) string {
	m := map[string]string{"go": "golang", "sh": "shell"}
	return m[s]
}