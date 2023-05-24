package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var time12 string
	log.Print("Masukkan waktu dalam format 12 jam (HH:MM:SS AM/PM): ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	time12 = scan.Text()

	time24, err := convertTime(time12)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Waktu dalam format 24 jam : ", time24)
}

func convertTime(time12 string) (string, error) {
	partsTime := strings.Split(time12, " ")
	if len(partsTime) != 2 {
		return "", errors.New("Invalid input format")
	}

	timeStr := partsTime[0]
	meridiem := strings.ToUpper(partsTime[1])

	if meridiem != "AM" && meridiem != "PM" {
		return "", errors.New("Invalid input format")
	}

	timeParts := strings.Split(timeStr, ":")
	if len(timeParts) != 3 {
		return "", errors.New("Invalid input format")
	}

	hourStr := timeParts[0]

	minute, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return "", errors.New("Invalid input format")
	}

	if minute >= 60 || minute < 0 {
		return "", errors.New("Invalid input format")
	}

	second, err := strconv.Atoi(timeParts[2])
	if err != nil {
		return "", errors.New("Invalid input format")
	}

	if second >= 60 || second < 0 {
		return "", errors.New("Invalid input format")
	}

	minuteStr := timeParts[1] + ":" + timeParts[2]

	hour, err := convertHour(hourStr, meridiem)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%02d:%s", hour, minuteStr), nil
}

func convertHour(hourStr string, meridiem string) (int, error) {
	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return 0, errors.New("Invalid input format")
	}

	if hour < 1 || hour > 12 {
		return 0, errors.New("Invalid input format")
	}

	switch meridiem {
	case "PM":
		if hour != 12 {
			hour += 12
		}
	case "AM":
		if hour == 12 {
			hour = 0
		}
	}

	return hour, nil
}
