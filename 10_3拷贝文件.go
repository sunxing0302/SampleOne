// 10_3拷贝文件.go
/*
将文件生成exe后，传入3个参数：
10_3拷贝文件.exe xxx.png yyy.png
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]

	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	//只读方式打开
	srcFile, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}

	//新建目的文件
	dstFile, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return
	}

	//操作完毕，需要关闭文件
	defer srcFile.Close()
	defer dstFile.Close()

	//核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 4*1024)
	for {
		n, err := srcFile.Read(buf) //从源文件读取内容
		//fmt.Println(n)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("读取源文件错误。err:", err)
		}

		//往目标文件里写，读多少写多少
		dstFile.Write(buf[:n])

	}

}
