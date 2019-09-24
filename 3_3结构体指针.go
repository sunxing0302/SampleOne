// Test2.go
package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func Zhizheng() {
	//1、指针有合法指向后，才操作成员
	//先定义一个普通结构体变量
	var s Student
	//在定义一个指针变量，保存s的地址
	var p1 *Student
	p1 = &s

	//通过指针操作成员  p1.id 和(*p1).id完全等价，只能使用.运算符
	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("p1 = ", p1)
	fmt.Println("s = ", s)

	//2、通过new申请一个结构体
	p2 := new(Student)
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 = ", p2)
}

func main_2() {
	fmt.Println("Hello World!")
	/*
		//初始化
		var s Student = Student{1, "sunxing", 'm', 19}
		fmt.Println(s)

		s2 := Student{2, "zhangsan", 'm', 20}
		fmt.Println(s2)

		s3 := Student{id: 3, name: "lishi"}
		fmt.Println(s3)

		var s4 Student
		s4.name = "wangwu"
		s4.age = 100
		s4.id = 7
		fmt.Println(s4) */
	Zhizheng()
}
