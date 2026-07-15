package utils

import "os"

func Clear() {
	os.Stdout.Write([]byte("\x1b[H\x1b[2J"))
}
