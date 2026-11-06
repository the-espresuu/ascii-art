package process

import "strings"

func ParseCharacters(data string) [][]string {
	data = strings.ReplaceAll(data, "\r", "")
	line := strings.Split(data, "\n")
	chars := make([][]string, 0, 95)

	for i := 0; i+7 < len(line); i += 9 {
		block := line[i+1 : i+9]
		chars = append(chars, block)
	}

	return chars
}
