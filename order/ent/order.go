// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sei-ri/microservice.io/order/ent/order"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges OrderEdges `json:"edges"`
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Items holds the value of the items edge.
	Items []*Item `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) ItemsOrErr() ([]*Item, error) {
	if e.loadedTypes[0] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldUserID, order.FieldStatus:
			values[i] = &sql.NullString{}
		case order.FieldCreatedAt, order.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				o.ID = value.String
			}
		case order.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				o.UserID = value.String
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = value.String
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryItems queries the "items" edge of the Order entity.
func (o *Order) QueryItems() *ItemQuery {
	return (&OrderClient{config: o.config}).QueryItems(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(o.UserID)
	builder.WriteString(", status=")
	builder.WriteString(o.Status)
	builder.WriteString(", created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}