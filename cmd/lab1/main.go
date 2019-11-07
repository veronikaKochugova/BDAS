package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	sourceAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	targetAlpha = "QAZWSXEDCRFVTGBYHNUJMIKOLP"
	sourceDigh  = "0123456789"
	targetDigh  = "2534678901"
)

var (
	source = sourceAlpha + strings.ToLower(sourceAlpha) + sourceDigh
	target = targetAlpha + strings.ToLower(targetAlpha) + targetDigh
)

func Obfuscate(s string) (result string) {
	for _, r := range s {
		si := strings.Index(source, string(r))
		if si != -1 {
			result += string(target[si])
		} else {
			result += string(r)
		}
	}
	return result
}

func Unobfuscate(s string) (result string) {
	for _, r := range s {
		// TODO specific symbols
		ti := strings.Index(target, string(r))
		sr := source[ti]
		result += string(sr)
	}
	return result
}

func main() {
	input, iErr := os.Open("examples/example.xml")
	output, oErr := os.OpenFile("examples/obfuscated-example.xml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)

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
			obfuscated := Obfuscate(line)
			fmt.Println(obfuscated)
			if _, wErr := writer.WriteString(obfuscated); wErr != nil {
				log.Fatal("Write error: ", wErr)
				break
			}
		} else {
			break
		}
	}
	writer.Flush()
}
