package ask

import (
	"bufio"
	"fmt"
    "os"
)
func MongoDBAskForTypeOfAction() {
	reader := bufio.NewReader(os.Stdin)

    // Pergunta ao usuário
    fmt.Println("\n\nWhich type of action would you do?")
    fmt.Println("1. Generate table and schema")
    fmt.Println("2. Remove table and schema")
    fmt.Println("3. Setup it")
    fmt.Print("Action type: ")

    choice, _ := reader.ReadString('\n')

    // Responde com a escolha
    switch choice {
    case "1\n":
        fmt.Println("Generate table and schema (not implemented)")
    case "2\n":
        fmt.Println("Remove table and schema (not implemented)")
    case "3\n":
        fmt.Println("Setup it (not implemented).")
    default:
		fmt.Println("\n\nNot a listed action. \n\n")
        MongoDBAskForTypeOfAction()
    }
}