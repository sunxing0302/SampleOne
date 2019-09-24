//将json对象解决到结构体中
package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {

	jsonBuf := `
	{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
}`

	var tmp IT
	err := json.Unmarshal([]byte(jsonBuf), &tmp)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("tmp: %+v", tmp)

}
