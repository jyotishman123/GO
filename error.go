package main

import "fmt"

func main() {
	x, err := GetUser()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Wrong")
	} else {
		fmt.Println(x)
		fmt.Printf("Correct")
	}
}

func GetUser() (string, error) {  // Note the order of return types
	x := 0
	if x == 0 {
		return "", fmt.Errorf("Can't send anything")
	} else {
		return "This is correct",nil
	}
}
