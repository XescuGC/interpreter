package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xescugc/interpreter/entities"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("calc> ")
		text, _ := reader.ReadString('\n')
		//fmt.Printf("%#v", text)
		i := entities.NewInterpreter(text)
		fmt.Println(i.Expr())
	}
}
