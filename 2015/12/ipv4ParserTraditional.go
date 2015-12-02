package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/rayman22201/goParserCombinate/parserCombinator"
)

var dot = parserCombinator.CharParser('.', "DOT")

var isInt = regexp.MustCompile(`^[0-9]+`)

func parseInt(input string) (parserCombinator.ParseNode, string) {
	numberStr := isInt.FindString(input)
	if numberStr == "" {
		return nil, input
	}
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return nil, input
	}

	return number, strings.Replace(input, numberStr, "", 1)
}

func subnet(input string) (parserCombinator.ParseNode, string) {
	result, output := parseInt(input)
	if result == nil || result.(int) > 255 {
		return nil, input
	}
	return result, output
}

func ipv4(input string) (parserCombinator.ParseNode, string) {
	processIPV4 := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		var ipAddress string
		for _, node := range nodes {
			switch t := node.(type) {
			case int:
				ipAddress += strconv.Itoa(t)
			case string:
				ipAddress += "."
			}
		}
		return ipAddress
	}
	return parserCombinator.And(processIPV4, subnet, dot, subnet, dot, subnet, dot, subnet)(input)
}
