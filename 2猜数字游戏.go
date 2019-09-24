/*
1. 生成一个四位的随机数
2. 获取数据数的每一个数字
3. 通过键盘输入4个数字，然后与随机数比较大小
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成一个随机数
func CreateNum(p *int) {
	var number int
	rand.Seed(time.Now().UnixNano())

	for {
		number = rand.Intn(9999)
		if number > 1000 {
			break
		}
	}
	*p = number

}

//取得切片的每个数字
func GetNum(s []int, num int) {
	s[0] = num / 1000       //取千位
	s[1] = num % 1000 / 100 //取百位
	s[2] = num % 100 / 10   //取十位
	s[3] = num % 10         //取个位
	//fmt.Println(s)

}

//猜数字游戏
func OnGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		fmt.Println("请输入一个四位数")
		fmt.Scan(&num)

		for {
			if num > 999 && num < 10000 {
				break
			}
			fmt.Println("您输入的数不符合要求")
		}
		GetNum(keySlice, num)
		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一点\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 { //4位都猜对了
			fmt.Println("全部猜对!!!")
			break //跳出循环
		}

	}

}

func main() {
	var randNum int
	CreateNum(&randNum)
	//fmt.Println("randNum: ", randNum)

	randSlice := make([]int, 4)
	//保存这个4位数的每一位
	GetNum(randSlice, randNum)
	//fmt.Println("randSlice = ", randSlice)

	OnGame(randSlice) //游戏
}
