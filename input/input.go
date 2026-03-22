package input

import (
	"bufio"
	"os"
	"pet/todolist"
	"strings"

	"github.com/k0kubun/pp"
)

func Input() *bufio.Scanner {
	var scanner = bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		pp.Print("Ошибка ввода")
	}
	//text := scanner.Text()
	return scanner
}

func Komanda() {
	for {
		pp.Println("Введите команду")
		scanner := Input()
		text := strings.Fields(scanner.Text())
		if len(text) == 0 {
			pp.Println(errorNil)
			newEvents(scanner.Text(), errorNil)
			continue
		}
		switch text[0] {
		case "add":
			if len(text) < 3 {
				pp.Println(errorStroka)
				newEvents(scanner.Text(), errorStroka)
				continue
			}
			error := todolist.Todolist(text)
			newEvents(scanner.Text(), error)
		case "list":
			error := todolist.Alltodolist()
			newEvents(scanner.Text(), error)
		case "del":
			if len(text) < 2 {
				pp.Println(errorStroka)
				newEvents(scanner.Text(), errorStroka)
				continue
			}
			error := todolist.Deletetodolist(text[1])
			newEvents(scanner.Text(), error)
		case "done":
			if len(text) < 2 {
				pp.Println(errorStroka)
				newEvents(scanner.Text(), errorStroka)
				continue
			}
			newEvents(scanner.Text(), todolist.TodolistDone(text[1]))
		case "help":
			Printhelp()
			newEvents(scanner.Text(), "")
		case "events":
			printEvents()
		case "exit":
			return
		default:
			pp.Println(errorStroka)
			newEvents(scanner.Text(), errorStroka)
		}

	}
}
