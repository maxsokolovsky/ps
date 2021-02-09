package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	CreateCmd    = "create"
	CancelCmd    = "cancel"
	GetStatusCmd = "status"
)

var addr = flag.String("addr", ":4000", "HTTP network address")
var host = flag.String("host", "localhost", "HTTP network host")

func main() {
	flag.Parse()
	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	cmd := os.Args[1]
	args := os.Args[2:]
	switch cmd {
	case CreateCmd:
		err = HandleCreateCmd(args)
	case CancelCmd:
		err = HandleCancelCmd(args)
	case GetStatusCmd:
		err = HandleGetStatusCmd(args)
	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
