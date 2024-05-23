package main

import "fmt"

type app struct {
	apps string
}

type messageToSend struct {
	phoneNumber int
	message string
	src app
}

func test(m messageToSend){
	 fmt.Println(m)
	 fmt.Println(m.src.apps)

}

func (app messageToSend) update(new string) string {
	 return new
}

func main(){
	 m:= messageToSend{
		  phoneNumber:456,
		  message:"hi",
		  src: app{
			  apps: "whatsapp",
		  },
	 }

// anonymous structs

	 myCar:= struct {
		Make string
		Model string
   } {
	   Make: "Ford",
	   Model: "Raptor",
   }

	 test(m)
	 fmt.Println(myCar)
	 fmt.Println("struct")
	 fmt.Println(m.update("new message"))
}