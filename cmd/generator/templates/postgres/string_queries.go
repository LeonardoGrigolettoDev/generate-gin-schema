package templates

import (
	"fmt"

	"github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/generator/model"
)

func GetQueryForCreationTableOnDB(tableName string, primaryKeyName string, primaryKeyType string, columns []model.Column) string {
	formatedColumns := ""
	for i := 0; i < len(columns); i++ {
		column := columns[i]
		columnStr := fmt.Sprintf("%s %s %s,", column.ColumnName, column.ColumnType, column.ColumnType)
		formatedColumns = formatedColumns + columnStr
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s %s PRIMARY KEY, %s created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);", tableName, primaryKeyName, primaryKeyType, formatedColumns)
	return query
}
