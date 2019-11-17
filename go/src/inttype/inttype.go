package main

import "fmt"

//8进制 16进制 的类型
func main() {
	var n1 = 0777
	var n2 = 0X765
	fmt.Printf("%v %T\n", n1, n1)

	fmt.Printf("%v %T\n", n2, n2)
}
