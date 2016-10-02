package main

import "fmt"
import "os/user"

func main() {
	myUser, _ := user.Current()
	fmt.Println(myUser.Username)
	fmt.Println(myUser.Name)
}

