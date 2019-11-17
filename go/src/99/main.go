package main

// 99 乘法表
import (
	"fmt"
	"strconv"
)

func chengFaBiao() {
	var table [9][9]string // 定义二维表
	var value string
	for i := 0; i < 9; i++ { //行
		for j := 0; j <= i; j++ { //列
			num := (j + 1) * (i + 1)
			if num < 10 && j >= 1 {
				value = " " + strconv.Itoa(num)
			} else {
				value = strconv.Itoa(num)
			}
			table[i][j] = strconv.Itoa(j+1) + "*" + strconv.Itoa(i+1) + "=" + value
			fmt.Printf("%s\t", table[i][j])
		}
		fmt.Print("\n")
	}
}

func chengFaBiao2() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func chengFaBiao3() {
	for i := 1; i <= 10; i++ {
		for j := 9; j > 0; j-- {
			fmt.Printf("%d*%d=%d\t", i, j, j*i)
		}
		fmt.Println()
	}
}

func main() {
	chengFaBiao()
	chengFaBiao2()
	chengFaBiao3()
}
