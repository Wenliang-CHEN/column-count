package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

// TODO: refactor
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var columnCounts = make(map[string]int)
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "NULL" {
			line = "{any default value}"
		}
		items := strings.Split(scanner.Text(), ",")
		for _, item := range items {
			if strings.Contains(item, "{any excluded}") {
				continue
			}

			key := strings.ToLower(strings.Split(strings.Trim(item, " \"\\N"),":")[0]) 

			if key == "" {
				continue
			}

			column, _ := columnCounts[key]
			columnCounts[key] = (column + 1)
		}
		lineCount += 1
	}

	var columns = make([]Column, 0)
	for key, val := range columnCounts {
		ratio := float32(val) / float32(lineCount)
		if ratio < 0.01 {
			continue
		}
		columns = append(columns, Column{name: key, ratio: ratio})
	}

	sort.Slice(columns, func(i, j int) bool {
		return columns[i].ratio >= columns[j].ratio
	})

	for _, column := range columns {
		fmt.Printf("|%v|%.2f%%|\n", column.name, column.ratio*100)
	}
}

type Column struct {
	name  string
	ratio float32
}
