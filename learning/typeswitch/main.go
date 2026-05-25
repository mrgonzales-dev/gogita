package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ------------------------------------------------
// We define a few "message" types (any struct works).
// ------------------------------------------------
type TextMsg struct {
	Content string
}

type NumberMsg struct {
	Value int
}

type QuitMsg struct{}

// processMessage receives a value of type 'any' (empty interface)
// and uses a type switch to figure out what's inside.
func processMessage(msg any) {
	switch typedMsg := msg.(type) {

	case TextMsg:
		fmt.Println("It's text:", typedMsg.Content)

	case NumberMsg:
		fmt.Println("It's a number:", typedMsg.Value)

	case QuitMsg:
		fmt.Println("Quitting!")
		return

	default:
		// Runs if none of the above match.
		fmt.Println("Unknown type!")
	}
}

// ------------------------------------------------
// Without the type switch, you'd have to do this:
// ------------------------------------------------
func processMessageLong(msg any) {
	if textMsg, castOk := msg.(TextMsg); castOk {
		fmt.Println("It's text:", textMsg.Content)
	} else if numberMsg, castOk := msg.(NumberMsg); castOk {
		fmt.Println("It's a number:", numberMsg.Value)
	} else if _, castOk := msg.(QuitMsg); castOk {
		fmt.Println("Quitting!")
		return
	} else {
		fmt.Println("Unknown type!")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text or a number (type 'quit' to exit):")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		lower := strings.ToLower(line)
		if lower == "quit" || lower == "exit" {
			processMessage(QuitMsg{})
			break
		}
		//what is Atoi?
		// answer: it converts a string to an int
		if num, err := strconv.Atoi(line); err == nil {
			processMessage(NumberMsg{Value: num})
		} else {
			processMessage(TextMsg{Content: line})
		}
	}
}
