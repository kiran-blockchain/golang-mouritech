package main

import "fmt"

type tank interface{
	Area() float64
	Volume() float64
}

type demo struct{
	radius float64
	height float64
}

func (d demo) Area() float64{
	return 2*d.radius*d.height
}

func (d demo) Volume() float64{
	return d.radius*d.height
}

func (d demo) xyz() float64{
	return d.height+d.radius
}
func main(){
 var t tank
 t = demo{10,20}
 
 fmt.Println(t.Area())
 fmt.Println(t.Volume())
}