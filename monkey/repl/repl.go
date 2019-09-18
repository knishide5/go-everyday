package repl

import {
	"bufio"
	"fmt"
	"io"
	"go-everyday/monkey/lexer"
	"go-everyday/monkey/token"
}

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := buffio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
