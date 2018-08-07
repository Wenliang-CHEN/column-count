package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var profileTimes = make(map[string]float64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\"")
		elements := strings.Split(line, "#")
		if len(elements)  < 2{
			continue
		}

		profileString := elements[1]
		profile := strings.Split(profileString, ":")

		key := profile[0]
		secs, _ := profileTimes[key]

		endpointTime, _ := strconv.ParseFloat(strings.Trim(profile[1], " "), 32)
		profileTimes[key] = (secs + endpointTime)
		fmt.Printf("%v\n",  profile[1])
	}

	var total float64
	for _, val := range profileTimes {
		total += val
	}

	for key, val := range profileTimes {
		fmt.Printf("|%v|%.4f|%.2f%%|\n", key, val, val / total * 100)
	}
}
