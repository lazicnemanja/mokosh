package main

import (
	"log"
	"os"

	"github.com/lazicnemanja/mokosh/internal/argsparser"
	"github.com/lazicnemanja/mokosh/pkg/importer"
)

func main() {
	args := os.Args[1:]

	arguments := argsparser.Parse(args, []string{"path", "name", "version"})

	if arguments["path"] != "" {
		_, err := os.Open(arguments["path"])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Provide path to .mwb file!")
	}

	importer.Run(arguments["path"], arguments["name"], arguments["version"])
}
