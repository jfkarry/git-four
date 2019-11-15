package main

import (
	"fmt"
	"time"
)

func main() {
	var n,q,w,e,r int64
	var m string
	fmt.Println("input")
	fmt.Scanln(&n)
	fmt.Scanln(&q)
	fmt.Scanln(&w)
	fmt.Scanln(&e)
	fmt.Scanln(&r)
	fmt.Scanln(&m)
	if n<3600||q<3600||w<3600||e<3600||r<3600{
		fmt.Println("input error")
	}else{
		if m=="result"{
			fmt.Println("input ok!")
			fmt.Println("input ok!")
			fmt.Println("input ok!")
			fmt.Println("input ok!")
			fmt.Println("input ok!")
			fmt.Println("the result  are:")
			fmt.Println(time.Unix(n,1))
			fmt.Println(time.Unix(q,1))
			fmt.Println(time.Unix(w,1))
			fmt.Println(time.Unix(e,1))
			fmt.Println(time.Unix(r,1))
		}
	}
}

