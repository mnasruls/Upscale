package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// read file
	numbers, err := ReadFile("test1/angka.txt")
	if err != nil {
		log.Println("Error read file : ", err)
		panic(err)
	}

	log.Println("Total angka pada file : ", len(numbers))

	// sum numbers
	sum, err := SumNum(numbers)
	if err != nil {
		log.Println("Error sum number : ", err)
		panic(err)
	}

	log.Println("Jumlah semua angka : ", sum)
}

func ReadFile(filename string) ([]string, error) {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dataFile := string(file)

	numbers := strings.Fields(dataFile)

	return numbers, nil
}

func SumNum(numbers []string) (int, error) {

	var result int
	// loop to sum numbers
	for _, numS := range numbers {
		num, err := strconv.Atoi(numS)
		if err != nil {
			return 0, err
		}

		result += num
	}

	return result, nil
}
