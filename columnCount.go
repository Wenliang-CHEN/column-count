package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: refactor
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("unable to read file")
	}
	defer file.Close()

	var columnCounts = make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ",")
		for _, item := range items {
			if item == "" {
				continue
			}

			key := strings.Trim(item, " ")
			column, _ := columnCounts[key]
			columnCounts[key] = (column + 1)
		}
	}

	for key, val := range columnCounts {
		fmt.Printf("%v: %v\n", key, val)
	}
}
