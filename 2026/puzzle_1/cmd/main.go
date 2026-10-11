package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Threshold         = 60 //temperature threshold in degrees
	TimeToHeatSeconds = 1  //time to heat 1 degree in seconds
	TimeToCoolSeconds = 5  //time to cool 1 degree in seconds
	Delimiter         = '\n'
)

func main() {
	println(process("input.txt"))
}

func process(path string) int {
	total := int(0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString(Delimiter)
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(i)
			if i > Threshold {
				//cool
				secondsToCool := (i - Threshold) * TimeToCoolSeconds
				total += secondsToCool
			} else if i < Threshold {
				//heat
				secondsToHeat := (Threshold - i) * TimeToHeatSeconds
				total += secondsToHeat
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return total
}
