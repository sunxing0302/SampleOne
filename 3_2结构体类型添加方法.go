// 1_结构体类型添加方法.go
package main

import (
	"fmt"
)

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//带有接收者的函数叫方法
func (tmp Person) PrintInfo() {
	fmt.Println("Person :", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) setInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	var p Person
	p.name = "sx"
	p.PrintInfo()

	var p2 Person
	(&p2).setInfo("sunxing", 'a', 20)
	p2.PrintInfo()
}
