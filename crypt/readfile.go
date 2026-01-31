package crypt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFromFile(location string) string{
	f, err := os.Open(location)
	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	buffer := strings.Builder{}
	for scanner.Scan(){
		fmt.Print(".") 
		buffer.WriteString(scanner.Text()) 
	}
	f.Close()
	return buffer.String()
}
