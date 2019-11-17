package main

import "fmt"
var x [2]int
func arraySum(x [2]int) int{
	x :=x
 sum :=0
 for _,v:= rang x{
	 sum +=v
 }
 return sum
}
func main(){
	fmt.Print(arraySum([2]int{1, 2}))
}