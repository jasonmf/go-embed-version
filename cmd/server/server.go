package main

import (
	"flag"

	"versserv"
	"versserv/cmd"
)

func main() {
	flag.Parse()
	if *cmd.FVersion {
		cmd.ShowVersion()
	}

	s, err := versserv.NewServer()
	if err != nil {
		panic("creating server: " + err.Error())
	}

	if err := s.ListenAndServe(); err != nil {
		panic("listening: " + err.Error())
	}
}
