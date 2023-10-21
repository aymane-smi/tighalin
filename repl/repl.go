package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aymane-smi/tighalin/lexer"
	"github.com/aymane-smi/tighalin/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, user string) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println("@" + user + " % " + PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		//%+v\n => print the struct of the tok var with the keys
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
