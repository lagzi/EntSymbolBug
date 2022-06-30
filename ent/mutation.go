// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/lagzi/EntSymbolBug/ent/predicate"
	"github.com/lagzi/EntSymbolBug/ent/workercontainedinformation"
	"github.com/lagzi/EntSymbolBug/ent/workernetworksettings"

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
	TypeWorkerContainedInformation = "WorkerContainedInformation"
	TypeWorkerNetworkSettings      = "WorkerNetworkSettings"
)

// WorkerContainedInformationMutation represents an operation that mutates the WorkerContainedInformation nodes in the graph.
type WorkerContainedInformationMutation struct {
	config
	op                      Op
	typ                     string
	id                      *int
	clearedFields           map[string]struct{}
	network_settings        *int
	clearednetwork_settings bool
	done                    bool
	oldValue                func(context.Context) (*WorkerContainedInformation, error)
	predicates              []predicate.WorkerContainedInformation
}

var _ ent.Mutation = (*WorkerContainedInformationMutation)(nil)

// workercontainedinformationOption allows management of the mutation configuration using functional options.
type workercontainedinformationOption func(*WorkerContainedInformationMutation)

// newWorkerContainedInformationMutation creates new mutation for the WorkerContainedInformation entity.
func newWorkerContainedInformationMutation(c config, op Op, opts ...workercontainedinformationOption) *WorkerContainedInformationMutation {
	m := &WorkerContainedInformationMutation{
		config:        c,
		op:            op,
		typ:           TypeWorkerContainedInformation,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWorkerContainedInformationID sets the ID field of the mutation.
func withWorkerContainedInformationID(id int) workercontainedinformationOption {
	return func(m *WorkerContainedInformationMutation) {
		var (
			err   error
			once  sync.Once
			value *WorkerContainedInformation
		)
		m.oldValue = func(ctx context.Context) (*WorkerContainedInformation, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().WorkerContainedInformation.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWorkerContainedInformation sets the old WorkerContainedInformation of the mutation.
func withWorkerContainedInformation(node *WorkerContainedInformation) workercontainedinformationOption {
	return func(m *WorkerContainedInformationMutation) {
		m.oldValue = func(context.Context) (*WorkerContainedInformation, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WorkerContainedInformationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WorkerContainedInformationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WorkerContainedInformationMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WorkerContainedInformationMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().WorkerContainedInformation.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetNetworkSettingsID sets the "network_settings" edge to the WorkerNetworkSettings entity by id.
func (m *WorkerContainedInformationMutation) SetNetworkSettingsID(id int) {
	m.network_settings = &id
}

// ClearNetworkSettings clears the "network_settings" edge to the WorkerNetworkSettings entity.
func (m *WorkerContainedInformationMutation) ClearNetworkSettings() {
	m.clearednetwork_settings = true
}

// NetworkSettingsCleared reports if the "network_settings" edge to the WorkerNetworkSettings entity was cleared.
func (m *WorkerContainedInformationMutation) NetworkSettingsCleared() bool {
	return m.clearednetwork_settings
}

// NetworkSettingsID returns the "network_settings" edge ID in the mutation.
func (m *WorkerContainedInformationMutation) NetworkSettingsID() (id int, exists bool) {
	if m.network_settings != nil {
		return *m.network_settings, true
	}
	return
}

// NetworkSettingsIDs returns the "network_settings" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// NetworkSettingsID instead. It exists only for internal usage by the builders.
func (m *WorkerContainedInformationMutation) NetworkSettingsIDs() (ids []int) {
	if id := m.network_settings; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetNetworkSettings resets all changes to the "network_settings" edge.
func (m *WorkerContainedInformationMutation) ResetNetworkSettings() {
	m.network_settings = nil
	m.clearednetwork_settings = false
}

// Where appends a list predicates to the WorkerContainedInformationMutation builder.
func (m *WorkerContainedInformationMutation) Where(ps ...predicate.WorkerContainedInformation) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *WorkerContainedInformationMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (WorkerContainedInformation).
func (m *WorkerContainedInformationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WorkerContainedInformationMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WorkerContainedInformationMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WorkerContainedInformationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown WorkerContainedInformation field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkerContainedInformationMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown WorkerContainedInformation field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WorkerContainedInformationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WorkerContainedInformationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkerContainedInformationMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown WorkerContainedInformation numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WorkerContainedInformationMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WorkerContainedInformationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WorkerContainedInformationMutation) ClearField(name string) error {
	return fmt.Errorf("unknown WorkerContainedInformation nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WorkerContainedInformationMutation) ResetField(name string) error {
	return fmt.Errorf("unknown WorkerContainedInformation field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WorkerContainedInformationMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.network_settings != nil {
		edges = append(edges, workercontainedinformation.EdgeNetworkSettings)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WorkerContainedInformationMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case workercontainedinformation.EdgeNetworkSettings:
		if id := m.network_settings; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WorkerContainedInformationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WorkerContainedInformationMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WorkerContainedInformationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearednetwork_settings {
		edges = append(edges, workercontainedinformation.EdgeNetworkSettings)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WorkerContainedInformationMutation) EdgeCleared(name string) bool {
	switch name {
	case workercontainedinformation.EdgeNetworkSettings:
		return m.clearednetwork_settings
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WorkerContainedInformationMutation) ClearEdge(name string) error {
	switch name {
	case workercontainedinformation.EdgeNetworkSettings:
		m.ClearNetworkSettings()
		return nil
	}
	return fmt.Errorf("unknown WorkerContainedInformation unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WorkerContainedInformationMutation) ResetEdge(name string) error {
	switch name {
	case workercontainedinformation.EdgeNetworkSettings:
		m.ResetNetworkSettings()
		return nil
	}
	return fmt.Errorf("unknown WorkerContainedInformation edge %s", name)
}

// WorkerNetworkSettingsMutation represents an operation that mutates the WorkerNetworkSettings nodes in the graph.
type WorkerNetworkSettingsMutation struct {
	config
	op                                  Op
	typ                                 string
	id                                  *int
	clearedFields                       map[string]struct{}
	worker_contained_information        *int
	clearedworker_contained_information bool
	done                                bool
	oldValue                            func(context.Context) (*WorkerNetworkSettings, error)
	predicates                          []predicate.WorkerNetworkSettings
}

var _ ent.Mutation = (*WorkerNetworkSettingsMutation)(nil)

// workernetworksettingsOption allows management of the mutation configuration using functional options.
type workernetworksettingsOption func(*WorkerNetworkSettingsMutation)

// newWorkerNetworkSettingsMutation creates new mutation for the WorkerNetworkSettings entity.
func newWorkerNetworkSettingsMutation(c config, op Op, opts ...workernetworksettingsOption) *WorkerNetworkSettingsMutation {
	m := &WorkerNetworkSettingsMutation{
		config:        c,
		op:            op,
		typ:           TypeWorkerNetworkSettings,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWorkerNetworkSettingsID sets the ID field of the mutation.
func withWorkerNetworkSettingsID(id int) workernetworksettingsOption {
	return func(m *WorkerNetworkSettingsMutation) {
		var (
			err   error
			once  sync.Once
			value *WorkerNetworkSettings
		)
		m.oldValue = func(ctx context.Context) (*WorkerNetworkSettings, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().WorkerNetworkSettings.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWorkerNetworkSettings sets the old WorkerNetworkSettings of the mutation.
func withWorkerNetworkSettings(node *WorkerNetworkSettings) workernetworksettingsOption {
	return func(m *WorkerNetworkSettingsMutation) {
		m.oldValue = func(context.Context) (*WorkerNetworkSettings, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WorkerNetworkSettingsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WorkerNetworkSettingsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WorkerNetworkSettingsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *WorkerNetworkSettingsMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().WorkerNetworkSettings.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetWorkerContainedInformationID sets the "worker_contained_information" edge to the WorkerContainedInformation entity by id.
func (m *WorkerNetworkSettingsMutation) SetWorkerContainedInformationID(id int) {
	m.worker_contained_information = &id
}

// ClearWorkerContainedInformation clears the "worker_contained_information" edge to the WorkerContainedInformation entity.
func (m *WorkerNetworkSettingsMutation) ClearWorkerContainedInformation() {
	m.clearedworker_contained_information = true
}

// WorkerContainedInformationCleared reports if the "worker_contained_information" edge to the WorkerContainedInformation entity was cleared.
func (m *WorkerNetworkSettingsMutation) WorkerContainedInformationCleared() bool {
	return m.clearedworker_contained_information
}

// WorkerContainedInformationID returns the "worker_contained_information" edge ID in the mutation.
func (m *WorkerNetworkSettingsMutation) WorkerContainedInformationID() (id int, exists bool) {
	if m.worker_contained_information != nil {
		return *m.worker_contained_information, true
	}
	return
}

// WorkerContainedInformationIDs returns the "worker_contained_information" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// WorkerContainedInformationID instead. It exists only for internal usage by the builders.
func (m *WorkerNetworkSettingsMutation) WorkerContainedInformationIDs() (ids []int) {
	if id := m.worker_contained_information; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetWorkerContainedInformation resets all changes to the "worker_contained_information" edge.
func (m *WorkerNetworkSettingsMutation) ResetWorkerContainedInformation() {
	m.worker_contained_information = nil
	m.clearedworker_contained_information = false
}

// Where appends a list predicates to the WorkerNetworkSettingsMutation builder.
func (m *WorkerNetworkSettingsMutation) Where(ps ...predicate.WorkerNetworkSettings) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *WorkerNetworkSettingsMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (WorkerNetworkSettings).
func (m *WorkerNetworkSettingsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WorkerNetworkSettingsMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WorkerNetworkSettingsMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WorkerNetworkSettingsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown WorkerNetworkSettings field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkerNetworkSettingsMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown WorkerNetworkSettings field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WorkerNetworkSettingsMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WorkerNetworkSettingsMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkerNetworkSettingsMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown WorkerNetworkSettings numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WorkerNetworkSettingsMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WorkerNetworkSettingsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WorkerNetworkSettingsMutation) ClearField(name string) error {
	return fmt.Errorf("unknown WorkerNetworkSettings nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WorkerNetworkSettingsMutation) ResetField(name string) error {
	return fmt.Errorf("unknown WorkerNetworkSettings field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WorkerNetworkSettingsMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.worker_contained_information != nil {
		edges = append(edges, workernetworksettings.EdgeWorkerContainedInformation)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WorkerNetworkSettingsMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case workernetworksettings.EdgeWorkerContainedInformation:
		if id := m.worker_contained_information; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WorkerNetworkSettingsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WorkerNetworkSettingsMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WorkerNetworkSettingsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedworker_contained_information {
		edges = append(edges, workernetworksettings.EdgeWorkerContainedInformation)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WorkerNetworkSettingsMutation) EdgeCleared(name string) bool {
	switch name {
	case workernetworksettings.EdgeWorkerContainedInformation:
		return m.clearedworker_contained_information
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WorkerNetworkSettingsMutation) ClearEdge(name string) error {
	switch name {
	case workernetworksettings.EdgeWorkerContainedInformation:
		m.ClearWorkerContainedInformation()
		return nil
	}
	return fmt.Errorf("unknown WorkerNetworkSettings unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WorkerNetworkSettingsMutation) ResetEdge(name string) error {
	switch name {
	case workernetworksettings.EdgeWorkerContainedInformation:
		m.ResetWorkerContainedInformation()
		return nil
	}
	return fmt.Errorf("unknown WorkerNetworkSettings edge %s", name)
}
