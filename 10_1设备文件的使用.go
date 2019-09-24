// 10_1设备文件的使用.go
package main

import (
	"fmt"
	"os"
)

func main() {

	//os.Stdout.Close() //关闭后，无法输出
	fmt.Println("fmt.Println=hihi")
	os.Stdout.Write([]byte("os.Stdout.Write==hihi \n")) //往标准输出设备(屏幕)写内容

	//标准设备文件(os.Stdout)默认已经打开，yoghurt可以直接使用

	os.Stdin.Close() //关闭后，无法输入
	var a int
	fmt.Println("请输入a：")
	fmt.Scan(&a)
	fmt.Println("a = ", a)

	os.Create()
	os.NewFile()

}
