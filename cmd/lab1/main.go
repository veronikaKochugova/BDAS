package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	lib "./obfuscation"
	"github.com/akamensky/argparse"
)

func parseArgs() (input, output string, deobfustactionFlag bool) {
	parser := argparse.NewParser("main.go", "Obfuscate/Deobfuscate text from input file")
	i := parser.String("i", "input", &argparse.Options{Required: true, Help: "Path to input file (required)"})
	o := parser.String("o", "output", &argparse.Options{Required: false, Help: "Path to output file (by default: [input file].result)"})
	d := parser.Flag("d", "deobfuscate", &argparse.Options{Required: false, Help: "Deobfuscation flag"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	// If an output is not specified, then an input with a suffix ".result" is used.
	if *o == "" {
		*o = *i + ".result"
	}
	return *i, *o, *d
}

func main() {
	i, o, d := parseArgs()
	input, iErr := os.Open(i)
	output, oErr := os.OpenFile(o, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)

	defer input.Close()
	defer output.Close()

	if iErr != nil || oErr != nil {
		log.Fatal(iErr, oErr)
		return
	}

	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)
	for {
		if line, rErr := reader.ReadString('\n'); rErr == nil {
			var newLine string
			if d {
				newLine = lib.Deobfuscate(line)
			} else {
				newLine = lib.Obfuscate(line)
			}
			if _, wErr := writer.WriteString(newLine); wErr != nil {
				log.Fatal("Write error: ", wErr)
				break
			}
		} else {
			break
		}
	}
	writer.Flush()
}
