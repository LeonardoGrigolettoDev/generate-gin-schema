package ask

import (
	"bufio"
	"fmt"
    "os"
)
func AskForTypeOfDB() {
	reader := bufio.NewReader(os.Stdin)

    // Pergunta ao usuário
    fmt.Println("Which DB would you do an action?")
    fmt.Println("1. PostgreSQL")
    fmt.Println("2. MongoDB")
    fmt.Print("DB type: ")

    choice, _ := reader.ReadString('\n')

    // Responde com a escolha
    switch choice {
    case "1\n":
        fmt.Println("PostgreSQL")
		PostgreSQLAskForTypeOfAction()
    case "2\n":
        fmt.Println("MongoDB (not implemented)")
		MongoDBAskForTypeOfAction()
    default:
		fmt.Println("\n\nNot a listed action. \n\n")
        AskForTypeOfDB()
    }
}