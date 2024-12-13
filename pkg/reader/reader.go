package reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadIntByColumnFile(input string) (outputs [][]int, err error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			s := strings.Fields(string(line))
			for len(outputs) < len(s) {
				outputs = append(outputs, []int{})
			}
			for i, element := range s {
				value, err := strconv.Atoi(element)
				if err != nil {
					return nil, err
				}
				outputs[i] = append(outputs[i], value)
			}
		}
	}
	return outputs, nil
}

func ReadIntByLineFile(input string) (outputs [][]int, err error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			s := strings.Fields(string(line))
			lineInt := []int{}
			for _, element := range s {
				value, err := strconv.Atoi(element)
				if err != nil {
					return nil, err
				}
				lineInt = append(lineInt, value)
			}
			outputs = append(outputs, lineInt)
		}
	}
	return outputs, nil
}

func ReadStringByLineFile(input string) (outputs []string, err error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			outputs = append(outputs, string(line))
		}
	}
	return outputs, nil
}

func ReadCharFile(input string) (outputs [][]rune, err error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			lineRune := []rune{}
			for _, runeValue := range string(line) {
				lineRune = append(lineRune, runeValue)
			}
			outputs = append(outputs, lineRune)
		}
	}
	return outputs, nil
}
