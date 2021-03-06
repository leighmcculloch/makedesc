package main

import (
	"flag"
	"fmt"
	"os"

	"4d63.com/makefile"
)

var version = "<not set>"

func main() {
	printHelp := flag.Bool("help", false, "print this help list")
	printVersion := flag.Bool("version", false, "print program version")
	filename := flag.String("file", "Makefile", "a Makefile")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Makedesc is a tool for describing makefiles.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  makedesc [-file=Makefile]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *printHelp {
		flag.Usage()
		return
	}

	if *printVersion {
		fmt.Fprintln(os.Stderr, version)
		return
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading Makefile: %s", err)
		flag.Usage()
		return
	}
	defer file.Close()

	m, err := makefile.Unmarshal(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing Makefile: %s", err)
		flag.Usage()
		return
	}

	for _, target := range m.Targets {
		fmt.Fprintln(os.Stdout, target.Name)
	}
}
