package main

import (
	"flag"
	"log"
	"os"

	"github.com/TaCkeR00/gradecalc/internal/calculator"
)

func main() {
	filePath := flag.String("file", "", "file path")
	isInteractive := flag.Bool("interactive", false, "is the grades from file of the user fill them")
	flag.Parse()

	if *filePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	modules, err := calculator.FromJSON(file)
	if err != nil {
		log.Fatal(err)
	}

	if *isInteractive {
		err := ReadGradesFromUser(modules)
		if err != nil {
			log.Fatal(err)
		}
	}

	modules.CalculateGrade()

	PrintResult(modules)
}
