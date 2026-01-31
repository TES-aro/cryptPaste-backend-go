package crypt

import (
	"bufio"
	"os"
	"strings"
)

func readInput() string {

    // Buffered input that splits input on lines.
    input := bufio.NewScanner(os.Stdin)

    // Buffered output.
    output := bufio.NewWriter(os.Stdout)
    strBuilder := strings.Builder{}


    // Scan until EOF (no more input).
    for input.Scan() {
        text := input.Text()
        strBuilder.WriteString(text) 

    }
    _ = output.Flush()
    str := strBuilder.String()
    if(str != ""){
	    return str
    }
    return ""
}
