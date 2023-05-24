package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	FizzBuzz()         // bagian 1
	Palindrom()        // bagian 2
	Factorial()        // bagian 3
	BiggestSmallest()  // bagian 4
	LettersAndDigits() // bagian 5
}

// bagian 1
func FizzBuzz() {
	var inputInt string
	log.Print("Masukkan angka : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	inputInt = scan.Text()

	input, err := strconv.Atoi(inputInt)
	if err != nil {
		log.Println("format bilangan salah")
	}

	for i := 1; i <= input; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			log.Println("FizzBuzz")
		case i%3 == 0:
			log.Println("Fizz")
		case i%5 == 0:
			log.Println("Buzz")
		default:
			log.Println(i)
		}
	}
}

// bagian 2
func Palindrom() {
	var inputText string
	log.Print("Masukkan kalimat/kata : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	inputText = scan.Text()
	newText := strings.ToLower(inputText)

	var isPalindrom = true
	for i := 0; i < len(newText)/2; i++ {
		if newText[i] != newText[len(newText)-(i+1)] {
			isPalindrom = false
			break
		}
	}

	if isPalindrom {
		log.Println("Kata/kalimat yang diinput :", inputText, "\nMerupakan palindrom")
	} else {
		log.Println("Kata/kalimat yang diinput :", inputText, "\nBukan merupakan palindrom")
	}
}

// bagian 3
func Factorial() {
	var num int
	var result = 1
	log.Print("Masukkan bilangan: ")
	fmt.Scanln(&num)

	for i := num; i >= 1; i-- {
		result *= i
	}
	log.Println("Faktorial dari", num, "adalah", result)
}

// bagian 5
func BiggestSmallest() {
	var inputSInt string
	log.Print("Masukkan angka - angka (gunakan spasi untuk memisahkan urutan bilangan) : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	inputSInt = scan.Text()

	inputInt := strings.Split(inputSInt, " ")

	var arrInt []int
	for _, value := range inputInt {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Println("format angka salah")
			panic(err)
		}

		arrInt = append(arrInt, num)
	}

	sort.Ints(arrInt)

	log.Println("Bilangan terbesar adalah : ", arrInt[len(arrInt)-1])
	log.Println("Bilangan terkecil adalah : ", arrInt[0])
}

// bagian 6
func LettersAndDigits() {
	var inputText string
	var letters, digits, elses int
	log.Print("Masukkan kalimat/kata : ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	inputText = scan.Text()

	for _, value := range inputText {
		switch {
		case unicode.IsLetter(value):
			letters++
		case unicode.IsDigit(value):
			digits++
		default:
			elses++
		}
	}

	log.Println("Jumlah angka : ", digits)
	log.Println("Jumlah huruf : ", letters)
	log.Println("Jumlah lain-lain : ", elses)
}
