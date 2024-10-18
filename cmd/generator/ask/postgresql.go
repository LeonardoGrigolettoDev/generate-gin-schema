package ask

import (
	"fmt"
	"strings"

	"github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/generator/model"
	templates "github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/generator/templates/postgres"
)

func PostgreSQLAskForTypeOfAction() {
	questionStr := "\n\nWhich type of action would you do?\n1. Generate table and schema\n2. Remove table and schema\n3. Setup it"
	choice := AskForSimpleQuestion(questionStr, 3, true)
	switch choice {
	case "1":
		fmt.Println("Generate table and schema")
		PostgreSQLAskForTableDetails()
	case "2":
		fmt.Println("Remove table and schema (not implemented)")
	case "3":
		fmt.Println("Setup it (not implemented).")
	default:
		fmt.Println("\nNot a listed action.")
		PostgreSQLAskForTypeOfAction()
	}
}

func PostgreSQLAskForTableDetails() {
	tableNameQuestion := "\n\nWhat is the name of the new DB table?"
	tableName := strings.ReplaceAll(AskForSimpleQuestion(tableNameQuestion, 0, false), "\n", "")
	primaryKeyNameQuestion := "\n\nWhat is the name of the column primary key?"
	primaryKeyName := strings.ReplaceAll(AskForSimpleQuestion(primaryKeyNameQuestion, 0, false), "\n", "")
	primaryKeyPKType := PostgreSQLAskForPrimaryKeyDetails()
	columnSchema := PostgreSQLAskForColumnSchema()
	queryStr := templates.GetQueryForCreationTableOnDB(tableName, primaryKeyName, primaryKeyPKType, columnSchema)
	fmt.Println("\n\nQuery generated:" + "\n" + queryStr)
}

func PostgreSQLAskForPrimaryKeyDetails() string {
	questionStr := "\n\nWhat is the ID type of the column primary key?\n1. Serial\n2. Bigserial\n3. UUID\n4. INTEGER\n5. Composite Keys"
	primaryKeyType := AskForSimpleQuestion(questionStr, 4, true)
	switch primaryKeyType {
	case "1":
		fmt.Println("Serial")
		primaryKeyType = "SERIAL"
	case "2":
		fmt.Println("Bigserial")
		primaryKeyType = "BIGSERIAL"
	case "3":
		fmt.Println("UUID")
		primaryKeyType = "UUID"
	case "4":
		fmt.Println("INTEGER")
		primaryKeyType = "INTEGER"
	//TODO IMPLEMENTAR AQUI
	// case "5\n":
	// 	fmt.Println("Composite Keys")
	// 	primaryKeyType = ""
	default:
		fmt.Println("\n\nNot a listed action.")
	}
	return primaryKeyType
}

func PostgreSQLAskForColumnSchema() []model.Column {
	var columns []model.Column
	columnTypeQuestion := `
	+----------------------+-----------------------+
	|      Category        |        Options        |
	+----------------------+-----------------------+
	| NUMBERS (1-6)        | 1. SMALLINT           |
	|                      | 2. INTEGER            |
	|                      | 3. BIGINT             |
	|                      | 4. DECIMAL/NUMERIC    |
	|                      | 5. REAL               |
	|                      | 6. DOUBLE PRECISION   |
	+----------------------+-----------------------+
	| TEXT (7-9)           | 7. CHAR               |
	|                      | 8. VARCHAR            |
	|                      | 9. TEXT               |
	+----------------------+-----------------------+
	| DATETIME (10-14)     | 10. DATE              |
	|                      | 11. TIME              |
	|                      | 12. TIMESTAMP         |
	|                      | 13. TIMESTAMPTZ       |
	|                      | 14. INTERVAL          |
	+----------------------+-----------------------+
	| BOOLEAN (15)         | 15. BOOLEAN           |
	+----------------------+-----------------------+
	| GEOMETRIC (16-22)    | 16. POINT             |
	|                      | 17. LINE              |
	|                      | 18. LSEG              |
	|                      | 19. BOX               |
	|                      | 20. PATH              |
	|                      | 21. POLYGON           |
	|                      | 22. CIRCLE            |
	+----------------------+-----------------------+
	| NET (23-25)          | 23. CIDR              |
	|                      | 24. INET              |
	|                      | 25. MACADDR           |
	+----------------------+-----------------------+
	| JSON (26-27)         | 26. JSON              |
	|                      | 27. JSONB             |
	+----------------------+-----------------------+
	| UUID (28)            | 28. UUID              |
	+----------------------+-----------------------+
	| ARRAY (29)           | 29. Array             |
	+----------------------+-----------------------+
	| INTERVAL (30-34)     | 30. DATERANGE         |
	|                      | 31. INT4RANGE         |
	|                      | 32. NUMRANGE          |
	|                      | 33. TSRANGE           |
	|                      | 34. TSTZRANGE         |
	+----------------------+-----------------------+
	`
	for {

		columnName := AskForSimpleQuestion("\n\nWhat is the name of the column (or type 'done' to finish)?", 0, false)
		columnName = strings.TrimSpace(columnName)
		if columnName == "done" {
			break
		}

		columnType := AskForSimpleQuestion(columnTypeQuestion, 34, true)
		switch columnType {
		case "1":
			fmt.Println("SMALLINT")
			columnType = "SMALLINT"

		case "2":
			fmt.Println("INTEGER")
			columnType = "INTEGER"

		case "3":
			fmt.Println("BIGINT")
			columnType = "BIGINT"

		case "4":
			fmt.Println("DECIMAL/NUMERIC")
			columnType = "DECIMAL/NUMERIC"

		case "5":
			fmt.Println("REAL")
			columnType = "REAL"

		case "6":
			fmt.Println("CHAR")
			columnType = "CHAR"

		case "7":
			fmt.Println("VARCHAR")
			columnType = "VARCHAR"

		case "8":
			fmt.Println("TEXT")
			columnType = "TEXT"

		case "9":
			fmt.Println("DATE")
			columnType = "DATE"

		case "10":
			fmt.Println("TIME")
			columnType = "TIME"

		case "11":
			fmt.Println("TIMESTAMP")
			columnType = "TIMESTAMP"

		case "12":
			fmt.Println("TIMESTAMPTZ")
			columnType = "TIMESTAMPTZ"

		case "13":
			fmt.Println("INTERVAL")
			columnType = "INTERVAL"

		case "14":
			fmt.Println("BOOLEAN")
			columnType = "BOOLEAN"

		case "15":
			fmt.Println("POINT")
			columnType = "POINT"

		case "16":
			fmt.Println("LINE")
			columnType = "LINE"

		case "17":
			fmt.Println("LSEG")
			columnType = "LSEG"

		case "18":
			fmt.Println("BOX")
			columnType = "BOX"

		case "19":
			fmt.Println("PATH")
			columnType = "PATH"

		case "20":
			fmt.Println("POLYGON")
			columnType = "POLYGON"

		case "21":
			fmt.Println("CIRCLE")
			columnType = "CIRCLE"

		case "22":
			fmt.Println("CIDR")
			columnType = "CIDR"

		case "23":
			fmt.Println("INET")
			columnType = "INET"

		case "24":
			fmt.Println("MACADDR")
			columnType = "MACADDR"

		case "25":
			fmt.Println("JSON")
			columnType = "JSON"

		case "26":
			fmt.Println("JSONB")
			columnType = "JSONB"

		case "27":
			fmt.Println("UUID")
			columnType = "UUID"

		case "28":
			fmt.Println("Array")
			columnType = "Array"

		case "29":
			fmt.Println("DATERANGE")
			columnType = "DATERANGE"

		case "30":
			fmt.Println("INT4RANGE")
			columnType = "INT4RANGE"

		case "31":
			fmt.Println("NUMRANGE")
			columnType = "NUMRANGE"

		case "32":
			fmt.Println("TSRANGE")
			columnType = "TSRANGE"

		case "33":
			fmt.Println("TSTZRANGE")
			columnType = "TSTZRANGE"

		default:
			fmt.Println("\n\nNot a listed action.")
		}
		columnType = strings.TrimSpace(columnType)
		formatedColumn := model.Column{
			ColumnName:    columnName,
			ColumnType:    columnType,
			ColumnOptions: "",
		}
		columns = append(columns, formatedColumn)
	}
	return columns
}

// CREATE TABLE users (
//     id SERIAL PRIMARY KEY,       -- ID do usuário, gerado automaticamente
//     username VARCHAR(50) NOT NULL, -- Nome de usuário, não pode ser nulo
//     email VARCHAR(100) NOT NULL UNIQUE, -- Email, deve ser único e não pode ser nulo
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Data de criação, padrão para a data atual
// );
