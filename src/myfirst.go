package  main

import "fmt"

var p,q int =100,200
var u,v  string

var arr[10] int

type mystruct struct {
	redius float64
	a int
}

type User struct {
	name string
	age int
	sex  string


}




func main() {
	//var c ,d int
	//c ,d = Foo(1,2)
	//fmt.Print
/*
	getValues := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getValues(100))

	ret :=myclodse(1,2)
	fmt.Println(ret(3,5))*/

	var  user User
	user.age = 100
	fmt.Println("结构体user：",user.age)

	var puser *User
	puser = &user
	puser.age=200
	fmt.Println(puser.age)



	alist :=[]int {1,2,5,57,2,5,45,45,4,1,36,7,1,1,9}


	aslice := alist[1:5]
	fmt.Println("切片",aslice)

	fmt.Println(alist)

	for i:=1;i<10;i++{
		arr[i]=i
	}
	for i:=1;i<8;i++{
		fmt.Printf("元素arr[%d]是：%d\n",i,arr[i])

	}


	ptr :=&alist
	fmt.Println(*ptr)



	var myst  mystruct
	myst.a=1
	myst.redius=2.0

	c := myst.myf()
	fmt.Println(c)

	j :=100
	var  ph *int = &j
	print(*ph)


}

func (m mystruct) myf() float64 {

	return 3.14 * m.redius

}



func Foo(a,b int) (int ,int) {
	return p,q
}

func myclodse(x,y int ) func(a,b int) int  {
		i := 1
		return func(a,b int) int {
			return i
		}


}

