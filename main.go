package main

import (
	"fmt"
	"dataBase"
)
func main(){
	fmt.Println("running main")
	a, err :=dataBase.ReadFile("test.file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a) 
	c, err := dataBase.ReadFile("no_file")
	if err != nil {
		fmt.Println((err))
	}
	fmt.Println(c) 

}
