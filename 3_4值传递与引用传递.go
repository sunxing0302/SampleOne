package main

import (
	"fmt"
)

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//修改成员变量的值
//接受者为普通变量，非指针，值语义，传递一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue &p =%p\n", &p)
	fmt.Println("SetInfoValue p=", p)
}

//接受者为指针变量，引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoPointer p = %p\n", p)
}

func main() {

	s1 := Person{"zhangsan", 'n', 18}
	fmt.Printf("s1 = %p\n", &s1) //打印地址

	//值语义 (值传递)
	//s1.SetInfoValue("lishi", 'm', 22)
	//fmt.Println("s1=", s1)

	//引用传递
	(&s1).SetInfoPointer("lishi", 'm', 22)
	fmt.Println("s1=", s1)

}
