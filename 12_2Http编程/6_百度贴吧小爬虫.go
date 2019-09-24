// 6_百度贴吧小爬虫.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("http.Get err1 = ", err1)
		return
	}

	defer resp.Body.Close()

	//读取网页Body内容
	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 { //读取结束，或者，出问题
			//fmt.Println("resp.Body.Read err2 = ", err2)
			return
		}
		result += string(buf[:n])
	}
	return
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//明确目标（要知道你准备再哪个范围或者网站去搜索）
	//http://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0  //下一页+50
	for i := start; i <= end; i++ {
		url := "http://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url = ", url)

		//2) 爬 (将所有的网站内容全部爬下来)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err =", err)
			continue
		}

		//将result写到文件中
		fileName := "c:\\temp\\" + strconv.Itoa(i) + ".html"
		f, err2 := os.Create(fileName)
		if err2 != nil {
			fmt.Println("os.Create err2 = ", err2)
			return
		}

		f.Write([]byte(result))

		f.Close()
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页(>=1) :")
	fmt.Scan(&start)
	fmt.Println("请输入终止页(>= 起始页) :")
	fmt.Scan(&end)

	DoWork(start, end)
}
