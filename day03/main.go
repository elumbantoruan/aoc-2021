package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	lines := strings.Split(text, "\n")
	power := powerConsumption(lines)
	fmt.Println(power)
}

func powerConsumption(lines []string) int64 {
	var (
		gamma   []byte
		epsilon []byte
	)
	var bits = lines[0]

	for i := 0; i < len(bits); i++ {
		var (
			count0Bits int
			count1Bits int
		)
		for _, line := range lines {

			if line[i] == '0' {
				count0Bits++
			} else {
				count1Bits++
			}

		}
		if count0Bits > count1Bits {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		} else {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		}
	}
	gammaDec, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonDec, _ := strconv.ParseInt(string(epsilon), 2, 64)

	return gammaDec * epsilonDec
}
