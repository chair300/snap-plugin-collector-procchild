package main

import (
	"os"

	"github.com/Staples-Inc/snap-plugin-collector-childprod/childproc"
	"github.com/intelsdi-x/snap/control/plugin"
)

func main() {
	plugin.Start(childproc.Meta(), childproc.New(), os.Args[1])
}
