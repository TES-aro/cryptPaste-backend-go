package crypt

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func saveToFile(bits []byte, ID string) {
	dir, err := os.Getwd()
	check(err)
  timeStamp := time.Now().Format(("15-36-44"))
  fmt.Printf("\n  time used:\n    %s\n", timeStamp)
  fileName := ID
  path1 := filepath.Join(dir, fileName)
  err = os.WriteFile(path1, bits, 0666)
  check(err)
}
