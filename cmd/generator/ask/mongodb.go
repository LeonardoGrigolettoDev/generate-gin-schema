package ask

import (
	"fmt"
)

func MongoDBAskForTypeOfAction() {
	questionStr := "\n\nWhich type of action would you do?\n1. Generate table and schema\n2. Remove table and schema\n3. Setup it"
	var choice string = AskForSimpleQuestion(questionStr, 3, true)
	switch choice {
	case "1":
		fmt.Println("Generate table and schema (not implemented)")
	case "2":
		fmt.Println("Remove table and schema (not implemented)")
	case "3":
		fmt.Println("Setup it (not implemented).")
	default:
		fmt.Println("\n\nNot a listed action.")
		MongoDBAskForTypeOfAction()
	}
}
