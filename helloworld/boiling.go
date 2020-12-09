package main

import "fmt"

const boilingF  =212.0

func main()  {
	var f =boilingF
	var c=(f-32)*5/9
	fmt.Printf("bolling point =%g duF or %g duF\n",f,c)
	
}
func fToC(f float64)float64{
	return (f-32)*5/9
}

func f()*int{
	v :=1
	return &v
}
func incr(p*int)int{
	*p++
	return *p
}