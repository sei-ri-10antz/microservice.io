// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sei-ri/microservice.io/order/ent/item"
	"github.com/sei-ri/microservice.io/order/ent/order"
	"github.com/sei-ri/microservice.io/order/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeItem  = "Item"
	TypeOrder = "Order"
)

// ItemMutation represents an operation that mutates the Item nodes in the graph.
type ItemMutation struct {
	config
	op              Op
	typ             string
	id              *int
	product_id      *int
	addproduct_id   *int
	qty             *int
	addqty          *int
	created_at      *time.Time
	updated_at      *time.Time
	clearedFields   map[string]struct{}
	order_id        map[string]struct{}
	removedorder_id map[string]struct{}
	clearedorder_id bool
	done            bool
	oldValue        func(context.Context) (*Item, error)
	predicates      []predicate.Item
}

var _ ent.Mutation = (*ItemMutation)(nil)

// itemOption allows management of the mutation configuration using functional options.
type itemOption func(*ItemMutation)

// newItemMutation creates new mutation for the Item entity.
func newItemMutation(c config, op Op, opts ...itemOption) *ItemMutation {
	m := &ItemMutation{
		config:        c,
		op:            op,
		typ:           TypeItem,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withItemID sets the ID field of the mutation.
func withItemID(id int) itemOption {
	return func(m *ItemMutation) {
		var (
			err   error
			once  sync.Once
			value *Item
		)
		m.oldValue = func(ctx context.Context) (*Item, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Item.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withItem sets the old Item of the mutation.
func withItem(node *Item) itemOption {
	return func(m *ItemMutation) {
		m.oldValue = func(context.Context) (*Item, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ItemMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ItemMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *ItemMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetProductID sets the "product_id" field.
func (m *ItemMutation) SetProductID(i int) {
	m.product_id = &i
	m.addproduct_id = nil
}

// ProductID returns the value of the "product_id" field in the mutation.
func (m *ItemMutation) ProductID() (r int, exists bool) {
	v := m.product_id
	if v == nil {
		return
	}
	return *v, true
}

// OldProductID returns the old "product_id" field's value of the Item entity.
// If the Item object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ItemMutation) OldProductID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldProductID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldProductID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProductID: %w", err)
	}
	return oldValue.ProductID, nil
}

// AddProductID adds i to the "product_id" field.
func (m *ItemMutation) AddProductID(i int) {
	if m.addproduct_id != nil {
		*m.addproduct_id += i
	} else {
		m.addproduct_id = &i
	}
}

// AddedProductID returns the value that was added to the "product_id" field in this mutation.
func (m *ItemMutation) AddedProductID() (r int, exists bool) {
	v := m.addproduct_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetProductID resets all changes to the "product_id" field.
func (m *ItemMutation) ResetProductID() {
	m.product_id = nil
	m.addproduct_id = nil
}

// SetQty sets the "qty" field.
func (m *ItemMutation) SetQty(i int) {
	m.qty = &i
	m.addqty = nil
}

// Qty returns the value of the "qty" field in the mutation.
func (m *ItemMutation) Qty() (r int, exists bool) {
	v := m.qty
	if v == nil {
		return
	}
	return *v, true
}

// OldQty returns the old "qty" field's value of the Item entity.
// If the Item object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ItemMutation) OldQty(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldQty is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldQty requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldQty: %w", err)
	}
	return oldValue.Qty, nil
}

// AddQty adds i to the "qty" field.
func (m *ItemMutation) AddQty(i int) {
	if m.addqty != nil {
		*m.addqty += i
	} else {
		m.addqty = &i
	}
}

// AddedQty returns the value that was added to the "qty" field in this mutation.
func (m *ItemMutation) AddedQty() (r int, exists bool) {
	v := m.addqty
	if v == nil {
		return
	}
	return *v, true
}

// ResetQty resets all changes to the "qty" field.
func (m *ItemMutation) ResetQty() {
	m.qty = nil
	m.addqty = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ItemMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ItemMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Item entity.
// If the Item object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ItemMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ItemMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ItemMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ItemMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Item entity.
// If the Item object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ItemMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ItemMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddOrderIDIDs adds the "order_id" edge to the Order entity by ids.
func (m *ItemMutation) AddOrderIDIDs(ids ...string) {
	if m.order_id == nil {
		m.order_id = make(map[string]struct{})
	}
	for i := range ids {
		m.order_id[ids[i]] = struct{}{}
	}
}

// ClearOrderID clears the "order_id" edge to the Order entity.
func (m *ItemMutation) ClearOrderID() {
	m.clearedorder_id = true
}

// OrderIDCleared returns if the "order_id" edge to the Order entity was cleared.
func (m *ItemMutation) OrderIDCleared() bool {
	return m.clearedorder_id
}

// RemoveOrderIDIDs removes the "order_id" edge to the Order entity by IDs.
func (m *ItemMutation) RemoveOrderIDIDs(ids ...string) {
	if m.removedorder_id == nil {
		m.removedorder_id = make(map[string]struct{})
	}
	for i := range ids {
		m.removedorder_id[ids[i]] = struct{}{}
	}
}

// RemovedOrderID returns the removed IDs of the "order_id" edge to the Order entity.
func (m *ItemMutation) RemovedOrderIDIDs() (ids []string) {
	for id := range m.removedorder_id {
		ids = append(ids, id)
	}
	return
}

// OrderIDIDs returns the "order_id" edge IDs in the mutation.
func (m *ItemMutation) OrderIDIDs() (ids []string) {
	for id := range m.order_id {
		ids = append(ids, id)
	}
	return
}

// ResetOrderID resets all changes to the "order_id" edge.
func (m *ItemMutation) ResetOrderID() {
	m.order_id = nil
	m.clearedorder_id = false
	m.removedorder_id = nil
}

// Op returns the operation name.
func (m *ItemMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Item).
func (m *ItemMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ItemMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.product_id != nil {
		fields = append(fields, item.FieldProductID)
	}
	if m.qty != nil {
		fields = append(fields, item.FieldQty)
	}
	if m.created_at != nil {
		fields = append(fields, item.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, item.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ItemMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case item.FieldProductID:
		return m.ProductID()
	case item.FieldQty:
		return m.Qty()
	case item.FieldCreatedAt:
		return m.CreatedAt()
	case item.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ItemMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case item.FieldProductID:
		return m.OldProductID(ctx)
	case item.FieldQty:
		return m.OldQty(ctx)
	case item.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case item.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Item field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ItemMutation) SetField(name string, value ent.Value) error {
	switch name {
	case item.FieldProductID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProductID(v)
		return nil
	case item.FieldQty:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetQty(v)
		return nil
	case item.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case item.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Item field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ItemMutation) AddedFields() []string {
	var fields []string
	if m.addproduct_id != nil {
		fields = append(fields, item.FieldProductID)
	}
	if m.addqty != nil {
		fields = append(fields, item.FieldQty)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ItemMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case item.FieldProductID:
		return m.AddedProductID()
	case item.FieldQty:
		return m.AddedQty()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ItemMutation) AddField(name string, value ent.Value) error {
	switch name {
	case item.FieldProductID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddProductID(v)
		return nil
	case item.FieldQty:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddQty(v)
		return nil
	}
	return fmt.Errorf("unknown Item numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ItemMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ItemMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ItemMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Item nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ItemMutation) ResetField(name string) error {
	switch name {
	case item.FieldProductID:
		m.ResetProductID()
		return nil
	case item.FieldQty:
		m.ResetQty()
		return nil
	case item.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case item.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Item field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ItemMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.order_id != nil {
		edges = append(edges, item.EdgeOrderID)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ItemMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case item.EdgeOrderID:
		ids := make([]ent.Value, 0, len(m.order_id))
		for id := range m.order_id {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ItemMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedorder_id != nil {
		edges = append(edges, item.EdgeOrderID)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ItemMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case item.EdgeOrderID:
		ids := make([]ent.Value, 0, len(m.removedorder_id))
		for id := range m.removedorder_id {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ItemMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedorder_id {
		edges = append(edges, item.EdgeOrderID)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ItemMutation) EdgeCleared(name string) bool {
	switch name {
	case item.EdgeOrderID:
		return m.clearedorder_id
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ItemMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Item unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ItemMutation) ResetEdge(name string) error {
	switch name {
	case item.EdgeOrderID:
		m.ResetOrderID()
		return nil
	}
	return fmt.Errorf("unknown Item edge %s", name)
}

// OrderMutation represents an operation that mutates the Order nodes in the graph.
type OrderMutation struct {
	config
	op            Op
	typ           string
	id            *string
	user_id       *string
	status        *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	items         map[int]struct{}
	removeditems  map[int]struct{}
	cleareditems  bool
	done          bool
	oldValue      func(context.Context) (*Order, error)
	predicates    []predicate.Order
}

var _ ent.Mutation = (*OrderMutation)(nil)

// orderOption allows management of the mutation configuration using functional options.
type orderOption func(*OrderMutation)

// newOrderMutation creates new mutation for the Order entity.
func newOrderMutation(c config, op Op, opts ...orderOption) *OrderMutation {
	m := &OrderMutation{
		config:        c,
		op:            op,
		typ:           TypeOrder,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOrderID sets the ID field of the mutation.
func withOrderID(id string) orderOption {
	return func(m *OrderMutation) {
		var (
			err   error
			once  sync.Once
			value *Order
		)
		m.oldValue = func(ctx context.Context) (*Order, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Order.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOrder sets the old Order of the mutation.
func withOrder(node *Order) orderOption {
	return func(m *OrderMutation) {
		m.oldValue = func(context.Context) (*Order, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OrderMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OrderMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Order entities.
func (m *OrderMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *OrderMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetUserID sets the "user_id" field.
func (m *OrderMutation) SetUserID(s string) {
	m.user_id = &s
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *OrderMutation) UserID() (r string, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldUserID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *OrderMutation) ResetUserID() {
	m.user_id = nil
}

// SetStatus sets the "status" field.
func (m *OrderMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *OrderMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldStatus(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *OrderMutation) ResetStatus() {
	m.status = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *OrderMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *OrderMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *OrderMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *OrderMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *OrderMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *OrderMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddItemIDs adds the "items" edge to the Item entity by ids.
func (m *OrderMutation) AddItemIDs(ids ...int) {
	if m.items == nil {
		m.items = make(map[int]struct{})
	}
	for i := range ids {
		m.items[ids[i]] = struct{}{}
	}
}

// ClearItems clears the "items" edge to the Item entity.
func (m *OrderMutation) ClearItems() {
	m.cleareditems = true
}

// ItemsCleared returns if the "items" edge to the Item entity was cleared.
func (m *OrderMutation) ItemsCleared() bool {
	return m.cleareditems
}

// RemoveItemIDs removes the "items" edge to the Item entity by IDs.
func (m *OrderMutation) RemoveItemIDs(ids ...int) {
	if m.removeditems == nil {
		m.removeditems = make(map[int]struct{})
	}
	for i := range ids {
		m.removeditems[ids[i]] = struct{}{}
	}
}

// RemovedItems returns the removed IDs of the "items" edge to the Item entity.
func (m *OrderMutation) RemovedItemsIDs() (ids []int) {
	for id := range m.removeditems {
		ids = append(ids, id)
	}
	return
}

// ItemsIDs returns the "items" edge IDs in the mutation.
func (m *OrderMutation) ItemsIDs() (ids []int) {
	for id := range m.items {
		ids = append(ids, id)
	}
	return
}

// ResetItems resets all changes to the "items" edge.
func (m *OrderMutation) ResetItems() {
	m.items = nil
	m.cleareditems = false
	m.removeditems = nil
}

// Op returns the operation name.
func (m *OrderMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Order).
func (m *OrderMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OrderMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.user_id != nil {
		fields = append(fields, order.FieldUserID)
	}
	if m.status != nil {
		fields = append(fields, order.FieldStatus)
	}
	if m.created_at != nil {
		fields = append(fields, order.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, order.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OrderMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case order.FieldUserID:
		return m.UserID()
	case order.FieldStatus:
		return m.Status()
	case order.FieldCreatedAt:
		return m.CreatedAt()
	case order.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OrderMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case order.FieldUserID:
		return m.OldUserID(ctx)
	case order.FieldStatus:
		return m.OldStatus(ctx)
	case order.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case order.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Order field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) SetField(name string, value ent.Value) error {
	switch name {
	case order.FieldUserID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case order.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case order.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case order.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OrderMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OrderMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Order numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OrderMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OrderMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OrderMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Order nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OrderMutation) ResetField(name string) error {
	switch name {
	case order.FieldUserID:
		m.ResetUserID()
		return nil
	case order.FieldStatus:
		m.ResetStatus()
		return nil
	case order.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case order.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OrderMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.items != nil {
		edges = append(edges, order.EdgeItems)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OrderMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case order.EdgeItems:
		ids := make([]ent.Value, 0, len(m.items))
		for id := range m.items {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OrderMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeditems != nil {
		edges = append(edges, order.EdgeItems)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OrderMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case order.EdgeItems:
		ids := make([]ent.Value, 0, len(m.removeditems))
		for id := range m.removeditems {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OrderMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareditems {
		edges = append(edges, order.EdgeItems)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OrderMutation) EdgeCleared(name string) bool {
	switch name {
	case order.EdgeItems:
		return m.cleareditems
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OrderMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Order unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OrderMutation) ResetEdge(name string) error {
	switch name {
	case order.EdgeItems:
		m.ResetItems()
		return nil
	}
	return fmt.Errorf("unknown Order edge %s", name)
}