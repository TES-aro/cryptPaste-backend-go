package dataBase

import(
	"bufio"
	"os"
	"strings"
)

func ReadFile(ID string) (string, error){
	file, err := os.Open(ID)
	if err != nil{
		return "", err
	}
	defer file.Close()
	str := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		str = append(str, scanner.Text())
	}
	return strings.Join(str[:], "\n"), nil
}
