package main

import (
	"fmt"
	"math"
	"os"
)
func main() {
	outputParams()
	 maxValue :=math.Max(1,2)
	 fmt.Println("1 and 2 max value:",maxValue)
	 fmt.Println("1+4 sum is:",sum(1,4))
}
func sum( a int, b int) int  {
	return a+b
}

func outputParams(){
	var s,sep string
	for i:=1; i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}
func test(){
	for true {
        break
	}
	for {
		break
	}
}