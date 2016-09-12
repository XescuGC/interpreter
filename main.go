package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("calc> ")
		text, _ := reader.ReadString('\n')
		//fmt.Printf("%#v", text)
		i := NewInterpreter(text)
		fmt.Println(i.expr())
	}
}
