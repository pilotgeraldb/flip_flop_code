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

var (
	AllTemps      []int
	Temps         []int
	ExpectedTemps []int
)

const (
	TimeToHeatSeconds = 1 //time to heat 1 degree in seconds
	TimeToCoolSeconds = 5 //time to cool 1 degree in seconds
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
			AllTemps = append(AllTemps, i)

		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	Temps = AllTemps[:(len(AllTemps) / 2)]
	ExpectedTemps = AllTemps[(len(AllTemps) / 2):]
	for idx, temp := range Temps {
		Threshold := ExpectedTemps[idx]
		if temp > Threshold {
			//cool
			secondsToCool := (temp - Threshold) * TimeToCoolSeconds
			total += secondsToCool
		} else if temp < Threshold {
			//heat
			secondsToHeat := (Threshold - temp) * TimeToHeatSeconds
			total += secondsToHeat
		}
	}
	return total
}
