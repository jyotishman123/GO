package main

import "fmt"

func main(){
   // loops in go

   for i:=0; i<10; i++ {
	 fmt.Println((i))
	 if i == 5 {
		return 
	 }
	 call()
   }
}

func call(){
	 fmt.Println("Loop")
}