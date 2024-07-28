// Code generated by ent, DO NOT EDIT.

package menuparam

import (
	"time"
)

const (
	// Label holds the string label denoting the menuparam type in the database.
	Label = "menu_param"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// Table holds the table name of the menuparam in the database.
	Table = "sys_menu_params"
	// MenusTable is the table that holds the menus relation/edge.
	MenusTable = "sys_menu_params"
	// MenusInverseTable is the table name for the Menu entity.
	// It exists in this package in order to avoid circular dependency with the "menu" package.
	MenusInverseTable = "sys_menus"
	// MenusColumn is the table column denoting the menus relation/edge.
	MenusColumn = "menu_params"
)

// Columns holds all SQL columns for menuparam fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldKey,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_menu_params"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"menu_params",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
