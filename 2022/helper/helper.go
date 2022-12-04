package helper

import (
	"bufio"
	"os"
)

func GetInput(path string) []string {
	file, err := os.Open(path + "/input.txt")
	arr := []string{}

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			panic(err)
		}
		arr = append(arr, scanner.Text())
	}
	return arr
}
