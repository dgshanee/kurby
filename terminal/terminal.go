package terminal

import (
	"os"

	"golang.org/x/term"
)

func GetTerminalSize() (height int, width int, err error) {
	h, w, err := term.GetSize(int(os.Stdout.Fd()))
	return h, w, err
}
