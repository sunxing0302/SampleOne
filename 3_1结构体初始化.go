// Test3.go
package main

import (
	"fmt"
)

//人
type Person struct {
	name string
	sex  byte
	age  int
}

//学生
type Student struct {
	Person // 匿名字段，那么默认Student就包含了Person的所有字段
	id     int
	addr   string
}

func main() {
	//顺序初始化
	s1 := Student{Person{"mike", 'm', 18}, 1, "sz"}
	//s1 = {Person:{name:mike sex:109 age:18} id:1 addr:sz}
	fmt.Printf("s1 = %+v\n", s1)

	//s2 := Student{"mike", 'm', 18, 1, "sz"} //err

	//部分成员初始化1
	s3 := Student{Person: Person{"lily", 'f', 19}, id: 2}
	//s3 = {Person:{name:lily sex:102 age:19} id:2 addr:}
	fmt.Printf("s3 = %+v\n", s3)

	//部分成员初始化2
	s4 := Student{Person: Person{name: "tom"}, id: 3}
	//s4 = {Person:{name:tom sex:0 age:0} id:3 addr:}
	fmt.Printf("s4 = %+v\n", s4)

	var s5 Student
	//s5.Person.name = "zhangsan"
	s5.addr = "add1"
	s5.name = "lishi"
	fmt.Printf("s5 = %+v\n", s5)

}
