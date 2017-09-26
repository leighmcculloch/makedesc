package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"4d63.com/makefile"
)

func main() {
	printHelp := flag.Bool("help", false, "print this help list")
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

	targetNames := []string{}
	for _, target := range m.Targets {
		targetNames = append(targetNames, target.Name)
	}

	fmt.Fprintf(os.Stderr, "Targets: %s\n", strings.Join(targetNames, ", "))
}
