// 6_1通过if实现类型断言.go
package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {

	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student{"mike", 666}

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值，data分别是i[0],i[1],i[2]
	for index, data := range i {
		//	fmt.Println(index, data)
		//_,ok := data.(int)
		//fmt.Println(ok)

		//第一个返回的是值，第二个返回判断结果的真假
		//data.(int)
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int，内容为 %d \n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string，内容为 %s \n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为struct，内容为 id=%d, name=%s \n", index, value.id, value.name)
		}

	}
}
