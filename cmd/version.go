package cmd

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version = "" // set at compile time with -ldflags "-X versserv/cmd.Version=x.y.yz"

	FVersion = flag.Bool("version", false, "show version and exit")
)

func ShowVersion() {
	fmt.Println("version:", Version)
	os.Exit(0)
}
