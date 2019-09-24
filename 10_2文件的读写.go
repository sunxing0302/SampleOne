// 10_2WriteString的使用.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//创建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//关闭文件
	defer f.Close()

	//文件写入数据
	for i := 0; i < 10; i++ {
		buf := fmt.Sprintf("i= %d \n", i)
		f.WriteString(buf)
	}

}

func ReadFile(path string) {
	//打开文件
	fin, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	//关闭文件
	defer fin.Close()

	//读取文件
	buf := make([]byte, 1024) //开辟1024个字节的slice作为缓冲
	n, err2 := fin.Read(buf)
	if err2 != nil && err2 != io.EOF { //文件出错，同时没有到结尾
		fmt.Println("读取文件失败:", err2)
		return
	}
	fmt.Println("n:", n)
	fmt.Println(string(buf)) //输出读取的内容

}

//按行读取文件
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	for {
		//遇到"\n"结束读取，但是"\n"读取进入
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err:", err)
		}

		fmt.Printf("buf: %s", string(buf))
	}

}

func main() {

	path := "./demo.txt"
	WriteFile(path)
	//ReadFile(path)
	//ReadFileLine(path)
}
