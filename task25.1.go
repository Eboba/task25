package main

import (
	"flag"
	"fmt"
	"strings"
)

func parseTest(text string, search string) bool {

	if strings.Contains(text, search) { // данная фукция отлично работает с строками UTF-8
		return true
	} else {
		return false
	}
}
func main() {

	fmt.Println("25.8 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Программа для нахождения подстроки в кириллической подстроке.")
	fmt.Println("------------")

	var (
		text   string
		search string
	)

	flag.StringVar(&text, "str", "", "Enter text")
	flag.StringVar(&search, "substr", "", "Search query")

	flag.Parse()

	fmt.Println(parseTest(text, search))
}
