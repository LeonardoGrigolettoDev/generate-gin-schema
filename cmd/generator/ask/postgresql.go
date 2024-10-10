package ask

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PostgreSQLAskForTypeOfAction() {
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
		fmt.Println("Generate table and schema")
		PostgreSQLAskForTableDetails()
	case "2\n":
		fmt.Println("Remove table and schema (not implemented)")
	case "3\n":
		fmt.Println("Setup it (not implemented).")
	default:
		fmt.Println("\nNot a listed action.")
		PostgreSQLAskForTypeOfAction()
	}
}

func PostgreSQLAskForTableDetails() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n\nWhat is the name of the new DB table?")
	tableName, _ := reader.ReadString('\n')

	fmt.Println("\n\nWhat is the name of the column primary key?")
	primaryKeyName, _ := reader.ReadString('\n')
	primaryKeyPKType := PostgreSQLAskForPrimaryKeyDetails(reader)
	columnSchema := PostgreSQLAskForColumnSchema(reader)
	fmt.Println(tableName, primaryKeyName, primaryKeyPKType, columnSchema)
}

func PostgreSQLAskForPrimaryKeyDetails(reader *bufio.Reader) string {
	fmt.Println("\n\nWhat is the ID type of the column primary key?")
	fmt.Println("1. Serial")
	fmt.Println("2. Bigserial")
	fmt.Println("3. UUID")
	fmt.Println("4. INTEGER")
	fmt.Println("5. Composite Keys")
	primaryKeyType, _ := reader.ReadString('\n')
	switch primaryKeyType {
	case "1\n":
		fmt.Println("Serial")
	case "2\n":
		fmt.Println("Bigserial")
	case "3\n":
		fmt.Println("UUID")
	case "4\n":
		fmt.Println("INTEGER")
	case "5\n":
		fmt.Println("Composite Keys")
	default:
		fmt.Println("\n\nNot a listed action.")
		PostgreSQLAskForPrimaryKeyDetails(reader)
	}
	return primaryKeyType
}

func PostgreSQLAskForColumnSchema(reader *bufio.Reader) [][]string {
	var columns [][]string

	for {
		fmt.Println("What is the name of the column (or type 'done' to finish)?")
		columnName, _ := reader.ReadString('\n')
		columnName = strings.TrimSpace(columnName)

		if columnName == "done" {
			break
		}

		fmt.Println("What is the type of the data to this column?")
		fmt.Println("\nNUMBERS (1-6)")
		fmt.Println("1. SMALLINT")
		fmt.Println("2. INTEGER")
		fmt.Println("3. BIGINT")
		fmt.Println("4. DECIMAL/NUMERIC")
		fmt.Println("5. REAL")
		fmt.Println("6. DOUBLE PRECISION")
		fmt.Println("\nTEXT (7-9)")
		fmt.Println("7. CHAR")
		fmt.Println("8. VARCHAR")
		fmt.Println("9. TEXT")
		fmt.Println("\nDATETIME (10-14)")
		fmt.Println("10. DATE")
		fmt.Println("11. TIME")
		fmt.Println("12. TIMESTAMP")
		fmt.Println("13. TIMESTAMPTZ")
		fmt.Println("14. INTERVAL")
		fmt.Println("\nBOOLEAN (15)")
		fmt.Println("15. BOOLEAN")
		fmt.Println("\nGEOMETRIC (16-22)")
		fmt.Println("16. POINT")
		fmt.Println("17. LINE")
		fmt.Println("18. LSEG")
		fmt.Println("19. BOX")
		fmt.Println("20 PATH")
		fmt.Println("21. POLYGON")
		fmt.Println("22. CIRCLE")
		fmt.Println("\nNET (23-25)")
		fmt.Println("23. CIDR")
		fmt.Println("24. INET")
		fmt.Println("25. MACADDR")
		fmt.Println("\nJSON (26-27)")
		fmt.Println("26. JSON")
		fmt.Println("27. JSONB")
		fmt.Println("\nUUID (28)")
		fmt.Println("28. UUID")
		fmt.Println("\nArray (29)")
		fmt.Println("29. Array")
		fmt.Println("\nInterval (30-34)")
		fmt.Println("30. DATERANGE")
		fmt.Println("31. INT4RANGE")
		fmt.Println("32. NUMRANGE")
		fmt.Println("33. TSRANGE")
		fmt.Println("34. TSTZRANGE")

		columnType, _ := reader.ReadString('\n')
		switch columnType {
		case "1\n":
			fmt.Println("Serial")
		case "2\n":
			fmt.Println("Bigserial")
		case "3\n":
			fmt.Println("UUID")
		case "4\n":
			fmt.Println("INTEGER")
		case "5\n":
			fmt.Println("Composite Keys")
		default:
			fmt.Println("\n\nNot a listed action.")
			PostgreSQLAskForPrimaryKeyDetails(reader)
		}
		columnType = strings.TrimSpace(columnType)

		// Adicionando a coluna à lista
		columns = append(columns, []string{columnName, columnType})
	}

	return columns
}

// CREATE TABLE users (
//     id SERIAL PRIMARY KEY,       -- ID do usuário, gerado automaticamente
//     username VARCHAR(50) NOT NULL, -- Nome de usuário, não pode ser nulo
//     email VARCHAR(100) NOT NULL UNIQUE, -- Email, deve ser único e não pode ser nulo
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Data de criação, padrão para a data atual
// );
