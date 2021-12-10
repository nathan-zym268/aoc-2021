package lib

import (
	"bufio"
	"log"
	"os"
)

func ReadData(file_name string) []string {
	var data []string
	f, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data

}
