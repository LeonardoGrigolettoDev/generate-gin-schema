package model

type Column struct {
	ColumnName    string
	ColumnType    string
	ColumnOptions string
}

type CreateTable struct {
	TableName      string
	PrimaryKeyName string
	PrimaryKeyType string
	Columns        []Column
}
