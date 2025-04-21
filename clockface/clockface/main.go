package main

import (
	"os"
	"time"

	"github.com/suic86/learn-go-with-tests/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Writer(os.Stdout, t)
}
