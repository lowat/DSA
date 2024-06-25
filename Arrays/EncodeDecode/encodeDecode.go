package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func encode(words []string) string {
	var buffer bytes.Buffer
	for _, word := range words {
		buffer.WriteString(fmt.Sprintf("%v", (word)))
		buffer.WriteString("#")
		buffer.WriteString(word)
	}
	return buffer.String()
}

func decode(word string) []string {
	res := []string{}
	i, left := 0, 0
	for i < len(word) {
		c := word[i]
		if c != '#' {
			i += 1
			continue
		}
		length, err := strconv.Atoi(word[left:i])
		if err != nil {
			print(err)
		}

		currWord := word[i+1 : i+length+1]
		res = append(res, currWord)
	}
	return res
}
