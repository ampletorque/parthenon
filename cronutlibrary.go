package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func crontabRead(crontabPath string) [][]string {
	f, err := os.Open(crontabPath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	// scanner := bufio.NewScanner(os.Stdin) // later implement stdin option from cli

	scanner := bufio.NewScanner(f)
	var lines [][]string
	for scanner.Scan() {
		rawLine := scanner.Text()
		if (!strings.HasPrefix(rawLine, "#")) && (rawLine != "") {
			line := regexp.MustCompile("\\s+").Split(rawLine, 6)
			lines = append(lines, line)
		}
	}
	return lines
}

func crontabParse(crontabRaw [][]string) []map[string]string {
	var crontabParsedListComplete []map[string]string

	for i := 0; i < len(crontabRaw); i++ {
		var crontabParsedListSingle map[string]string
		crontabParsedListSingle = make(map[string]string)

		min := crontabRaw[i][0]
		hr := crontabRaw[i][1]
		dom := crontabRaw[i][2]
		mon := crontabRaw[i][3]
		dow := crontabRaw[i][4]
		cmd := crontabRaw[i][5]

		crontabParsedListSingle["min"] = min
		crontabParsedListSingle["hr"] = hr
		crontabParsedListSingle["dom"] = dom
		crontabParsedListSingle["mon"] = mon
		crontabParsedListSingle["dow"] = dow
		crontabParsedListSingle["cmd"] = cmd

		crontabParsedListComplete = append(crontabParsedListComplete, crontabParsedListSingle)
	}

	fmt.Println(crontabParsedListComplete)
	return crontabParsedListComplete
}

func crontabGetNext(crontabParsedList []map[string]string) []map[string]string {
	var crontabNext []map[string]string

	return crontabNext
}
