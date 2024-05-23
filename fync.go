package main

import "fmt"

func st(s1 int, s2 int) int {
	 return s1 + s2
}

// func st(s1, s2 int) int {
// 	return s1 + s2
// }

func main(){

      x:=5
	  fmt.Println(inc(x))
	  fmt.Println(x)

	  fmt.Println("Function")
      fmt.Println(st(4,2))
	  fmt.Println(getChord())
}

func inc(x int) int {
     x = 10
	return x
}

func getChord()(x,y int){	 
	return 
}