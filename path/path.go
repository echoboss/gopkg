package main


import (
	"fmt"
	"path"
	"path/filepath"
	"os"
)

func main() {

	fmt.Println(path.Base(""),path.Base("/user/root"),path.Base("//////"))

	fmt.Println(path.Clean("/tooy/.3T/.robo-3t/.1.2.1/cache/"))
	fmt.Println(path.Clean("/users/../usr///local"))


	fmt.Println(path.Dir(""),path.Dir("/tooy/.3T/.robo-3t/.1.2.1/cache/contents.txt"))

	fmt.Println(path.Ext("/tooy/.3T/.robo-3t/.1.2.1/cache/contents.txt"),path.Ext("tooy/.3T/.robo-3t/.1.2.1"))

	fmt.Println(path.IsAbs("/tooy/.3T/.robo-3t/.1.2.1"))
	fmt.Println(path.Join("/opt", "..//deploy"))

	fmt.Println(path.Match("\\root","root"))
	fmt.Println(path.Match("\\.bash_profile", "root/.bash_profile"))
	fmt.Println(path.Split("https://pic4.zhimg.com/80/v2-9b1fadc0b82995605b31830e701df29b_hd.jpg"))

	fmt.Println(filepath.Rel("/root","/root/"))
	filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		fmt.Println("walk", path)
		fmt.Println("walk", info.Size())
		return nil
	})

}
