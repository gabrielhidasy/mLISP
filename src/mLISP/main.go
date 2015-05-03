package main

import (
	"fmt"
	"io"
	"github.com/bobappleyard/readline"
)


func main() {
	//runtime.GOMAXPROCS(2) //se nao usa uma so
	fmt.Println("GoHLisp V0.0.1, 2-clause BSD")
	prompt := "MyLisp > "
	for {
		l, err := readline.String(prompt)
		if err == io.EOF {
			fmt.Println("Bye Bye")
			break
		}
		if err != nil {
			fmt.Println("Error: ",err);
			break
		}
		readline.AddHistory(l)
		AST, err := parse_and_lex(l)
		if err != nil {
			fmt.Println("Error in parsin",err)
		}
		value, err := eval(AST)
		if err != nil {
			fmt.Println("Error in eval",err)
		}
		fmt.Println("V: ",value)
	}
}
