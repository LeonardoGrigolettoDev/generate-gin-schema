package ask

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AskForTypeOfDB() {
	questionStr := "\n\nWhich DB would you do an action?\n1. PostgreSQL\n2. MongoDB"
	choice := AskForSimpleQuestion(questionStr, 2, true)
	switch choice {
	case "1":
		fmt.Println("PostgreSQL")
		PostgreSQLAskForTypeOfAction()
	case "2":
		fmt.Println("MongoDB (not implemented)")
		MongoDBAskForTypeOfAction()
	default:
		fmt.Println("\n\nNot a listed action.")
		AskForTypeOfDB()
	}
}

func AskForSimpleQuestion(question string, optionsLength int, verify bool) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question)
	choice, _ := reader.ReadString('\n')
	if verify {
		choice = strings.ReplaceAll(choice, "\n", "")
		choiceAsNumber, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Println("\nInvalid choice, please select a valid one:")
			return AskForSimpleQuestion(question, optionsLength, verify)
		}
		if choiceAsNumber > optionsLength {
			fmt.Println("\nInvalid choice, please select a valid one:")
			return AskForSimpleQuestion(question, optionsLength, verify)
		} else {
			return choice
		}
	} else {
		return choice
	}
}
