package main

import (
	"fmt"
	"os"
)
//var变量声明可以只声明不初始化， 位置不限
//简短变量的生命必须在函数内部 := 形式是简短变量声明

func main()  {
	var a  int
	fmt.Println(a)//默认是0

	var i,j,k =1,2,5

	fmt.Println(i,j,k)

	p,q,t :=4,5,6
	fmt.Println(p,q,t)

	var f,err = os.Open("/Users/pig/go/src/awesomeProject/src/myfirst.go")
	if err == nil {
		fmt.Println(f.Name())
	}


	g := i//简短变量
	fmt.Println(g)

	

}
