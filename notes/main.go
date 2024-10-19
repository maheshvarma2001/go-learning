package main

import (
	"bufio"
	"course/note/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title := getUserInput("Input title: ")

	content := getUserInput("Input content: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	userNote.Save()

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")

	text = strings.TrimSuffix(text, "\r")

	return text
}
