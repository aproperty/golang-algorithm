package main

import (
	"fmt"
	"time"
)

func printStr(s string) {

	ch1 := make(chan int)

	go func() {
		for i := 0; i < len(s); i++ {
			ch1 <- i
			if i%2 == 0 {
				fmt.Println("g1:", string(s[i]))
			}
		}
	}()

	go func() {
		for i := 0; i < len(s); i++ {
			<-ch1
			if i%2 != 0 {
				fmt.Println("g2:", string(s[i]))
			}
		}
	}()

}

// 实现一个函数，功能：接收一个字符串参数，使用两个协程 按顺序 交替 输入 每个字符。
// 思路：利用 无缓冲通道 的 阻塞特性
func main() {
	printStr("hello,world")
	time.Sleep(time.Second)
}
