package repl

import (
	"bufio"
	"fmt"
	"monkey/lexer"
	"monkey/token"
	"os"
)

const PROMPT = ">>"

func StartRepl() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Printf(PROMPT)
		input, _ := reader.ReadString('\n')
		l := lexer.New(input)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
