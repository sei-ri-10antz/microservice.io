// Code generated by entc, DO NOT EDIT.

package item

import (
	"time"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldQty holds the string denoting the qty field in the database.
	FieldQty = "qty"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeOrderID holds the string denoting the order_id edge name in mutations.
	EdgeOrderID = "order_id"
	// Table holds the table name of the item in the database.
	Table = "items"
	// OrderIDTable is the table the holds the order_id relation/edge. The primary key declared below.
	OrderIDTable = "order_items"
	// OrderIDInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderIDInverseTable = "orders"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldProductID,
	FieldQty,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// OrderIDPrimaryKey and OrderIDColumn2 are the table columns denoting the
	// primary key for the order_id relation (M2M).
	OrderIDPrimaryKey = []string{"order_id", "item_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultQty holds the default value on creation for the "qty" field.
	DefaultQty int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)