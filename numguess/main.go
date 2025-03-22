package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("规则：")
	fmt.Println("1.计算机将在1到100之间随机选取一个数字,你需要猜测这个数字。")
	fmt.Println("2.你可以选择难度（简单、中等、困难）对应这不同的猜测机会。")
	fmt.Println("3.请输入你的猜测")
	fmt.Println("")
	var n int
	var i int
	var c int
	for {
		fmt.Println("请选择难度级别 (1:简单, 2:中等, 3:困难): ")
		fmt.Println("1.简单(10次机会)")
		fmt.Println("2.中等(5次机会)")
		fmt.Println("3.困难(3次机会)")
		fmt.Scan(&n)
		randomRange := rand.Intn(100)
		switch n {
		case 1:
			start := time.Now()
			for i = 0; i < 10; i++ {
				fmt.Printf("第%d次机会,请输入你猜测的数字：\n", i+1)
				j := 0
				fmt.Scan(&j)
				if j == randomRange {
					fmt.Printf("恭喜你，猜对了！你在%d次猜中\n", i+1)
					break
				} else if j > randomRange {
					fmt.Println("你猜的数字大了！")
				} else {
					fmt.Println("你猜的数字小了！")
				}

			}
			fmt.Printf("最终数字为：%d\n", randomRange)
			tt := time.Since(start)
			fmt.Printf("游戏结束，用时：%v\n", tt)
			fmt.Println("游戏结束！可重新开始游戏")
		case 2:
			start := time.Now()
			for i = 0; i < 5; i++ {
				fmt.Printf("第%d次机会,请输入你猜测的数字：\n", i+1)
				j := 0
				fmt.Scan(&j)
				if j == randomRange {
					fmt.Printf("恭喜你，猜对了！你在%d次猜中\n", i+1)
					break
				} else if j > randomRange {
					fmt.Println("你猜的数字大了！")
				} else {
					fmt.Println("你猜的数字小了！")
				}

			}
			fmt.Printf("最终数字为：%d\n", randomRange)
			tt := time.Since(start)
			fmt.Printf("游戏结束，用时：%v\n", tt)
			fmt.Println("游戏结束！可重新开始游戏")
		case 3:
			start := time.Now()
			for i = 0; i < 3; i++ {
				fmt.Printf("第%d次机会,请输入你猜测的数字：\n", i+1)
				j := 0
				fmt.Scan(&j)
				if j == randomRange {
					fmt.Printf("恭喜你，猜对了！你在%d次猜中\n", i+1)
					break
				} else if j > randomRange {
					fmt.Println("你猜的数字大了！")
				} else {
					fmt.Println("你猜的数字小了！")
				}

			}
			fmt.Printf("最终数字为：%d\n", randomRange)
			tt := time.Since(start)
			fmt.Printf("游戏结束，用时：%v\n", tt)
			fmt.Println("游戏结束！可重新开始游戏")
		}
		fmt.Println("重新开始有游戏输入1")
		fmt.Println("退出游戏输入0")
		fmt.Scan(&c)
		if c == 0 {
			break
		}

	}

}
