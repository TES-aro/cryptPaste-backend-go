package dataBase

import (
	"fmt"
	"os"
)

func WriteFile(content, ID string) error{
	file, err := os.Create(ID)
	if err != nil{
		fmt.Println(err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
