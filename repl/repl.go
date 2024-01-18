package repl

import (
	"bufio"
	"fmt"
	"github.com/K1ngArtes/monkey-lang-interpreter-go/lexer"
	"github.com/K1ngArtes/monkey-lang-interpreter-go/token"
	"io"
)

const PROMPT = ">> "

// Start starts REPL (Read Eval Print Loop) of accepting the user commands and printing back a set tokens doe this command
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
