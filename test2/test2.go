package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	var inputText string
	log.Print("Masukkan kalimat/kata : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	inputText = scan.Text()
	if CheckIfPolindrom(inputText) {
		log.Println("Kata/kalimat yang diinput :", inputText, "\nMerupakan palindrom")
	} else {
		log.Println("Kata/kalimat yang diinput :", inputText, "\nBukan merupakan palindrom")
	}

}

func CheckIfPolindrom(text string) bool {
	newText := strings.ToLower(text)

	for i := 0; i < len(newText)/2; i++ {
		if newText[i] != newText[len(newText)-(i+1)] {
			return false
		}
	}
	return true
}
