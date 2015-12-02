package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// @see http://blog.golang.org/laws-of-reflection
type ParseNode interface{}

type Parser func(input io.RuneScanner) bool

func parseChar(search rune, terminal string, outputChan chan parseNode) Parser {
	return func(input io.RuneScanner) bool {
		r, _, err := input.RuneReader.ReadRune()
		if err != nil {
			input.UnreadRune()
			outputChan <- err
			return false
		}
		if r == search {
			outputChan <- terminal
			return true
		} else {
			input.UnreadRune()
			return false
		}
	}
}

func parseNum(input io.RuneScanner) bool {
	r, _, err := input.RuneReader.ReadRune()
	if err != nil {
		outputChan <- err
	}
	if r < 48 || r > 57 {
		input.UnreadRune()
		return false
	} else {
		outputChan <- r
		return true
	}
}

func subnetParser(input io.RuneScanner) bool {
	numCount := 0
	for parseNum(input) == true {
		numCount++
	}
	if numCount < 1 || numCount > 3 {
		for i := 0; i < numCount; i++ {
			input.UnreadRune()
		}
		return false
	}
	outputChan <- numCount
	outputChan <- "SUBNET"
	return true
}

var outputChan = make(chan parseNode)

var dot = parseChar('.', "DOT", outputChan)

func ipv4Parser(input io.RuneScanner) bool {
	match := And(outputChan, subnetParser, dot, subnetParser, dot, subnetParser, dot, subnetParser)
	close(outputChan)
	return match
}

func ipv4Listener(outputChan chan parseNode) (string, error) {
	var results []parseNode
	for node := range outputChan {
		switch t := node.(type) {
		default:
			results = append(results, node)
		case error:
			return "", t
		case rune:
			// If it's a rune, convert it to a string
			results = append(results, fmt.Sprintf("%c", node))
		case string:
			if t == "SUBNET" {
				// Pop the subnet chars off the stack
				count, results = results[len(results)-1], results[:len(results)-1]
				subnetStr, results = strings.Join(results[len(results)-count:len(results)-1], ""), results[:len(results)-count]
				// Make sure it's a valid subnet
				number, err := strconv.Atoi(subnetStr)
				if err != nil {
					return nil, err
				}
				if number > 255 {
					return "", errors.New(fmt.Sprint("Invalid Subnet:", number))
				}
				// Push the subnet string back onto the stack
				results = append(results, subnetStr)
			}
		}
	}
	return strings.Join(results, sep), nil
}
