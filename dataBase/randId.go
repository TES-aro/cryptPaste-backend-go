package dataBase

import (
	"fmt"
	"math/rand"
)

func randNum() int {
	setNum := rand.Intn(3)
	if setNum == 0{
		return rand.Intn(11) + 48
	}
	if setNum == 1{
		return rand.Intn(26) + 65
	}
	return rand.Intn(26) + 97
}

func randChar() rune {
	return rune(randChar())
}

func GenID() string{
	runeArray := []rune{}
	for i := 0; i < 32; i++{
		runeArray = append(runeArray, randChar())
	}
	id := string(runeArray)
	fmt.Println("new id: ", id)
	return id
}
