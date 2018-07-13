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
		line := strings.Trim(scanner.Text(), "\"")
		if strings.ToLower(line) == "null" {
			line = "last_name,first_name,email,main-salary"
		}
		items := unique(strings.Split(line, ","))

		for _, item := range items {
			if strings.Contains(item, "dynamic_") {
				continue
			}

			key := strings.ToLower(strings.Split(item,":")[0]) 

			if key == ""{
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

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{} 
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}    
	return list
}

type Column struct {
	name  string
	ratio float32
}
