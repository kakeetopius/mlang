// Package repl contains the repl for the mlang language
package repl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/chzyer/readline"
	"github.com/kakeetopius/mlang/lexer"
	"github.com/kakeetopius/mlang/token"
	"github.com/pterm/pterm"
)

const PROMPT = ">>>"

func Start(in io.ReadCloser, out io.WriteCloser) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	rl, err := readline.NewEx(&readline.Config{
		Prompt:      pterm.FgLightBlue.Sprintf("%s%s ", "mlang", PROMPT),
		HistoryFile: filepath.Join(home, ".mlang"),
		Stdin:       in,
		Stdout:      out,
	})
	if err != nil {
		return err
	}

	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			return err
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
