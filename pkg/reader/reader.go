package reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadColumnIntInput(input string) (outputs [][]int, err error) {
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
