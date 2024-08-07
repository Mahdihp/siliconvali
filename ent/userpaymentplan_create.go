// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"siliconvali/ent/plan"
	"siliconvali/ent/user"
	"siliconvali/ent/userpaymentplan"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPaymentPlanCreate is the builder for creating a UserPaymentPlan entity.
type UserPaymentPlanCreate struct {
	config
	mutation *UserPaymentPlanMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (uppc *UserPaymentPlanCreate) SetAmount(i int64) *UserPaymentPlanCreate {
	uppc.mutation.SetAmount(i)
	return uppc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (uppc *UserPaymentPlanCreate) SetNillableAmount(i *int64) *UserPaymentPlanCreate {
	if i != nil {
		uppc.SetAmount(*i)
	}
	return uppc
}

// SetReferenceNumber sets the "reference_number" field.
func (uppc *UserPaymentPlanCreate) SetReferenceNumber(s string) *UserPaymentPlanCreate {
	uppc.mutation.SetReferenceNumber(s)
	return uppc
}

// SetTransactionNumber sets the "transaction_number" field.
func (uppc *UserPaymentPlanCreate) SetTransactionNumber(s string) *UserPaymentPlanCreate {
	uppc.mutation.SetTransactionNumber(s)
	return uppc
}

// SetNillableTransactionNumber sets the "transaction_number" field if the given value is not nil.
func (uppc *UserPaymentPlanCreate) SetNillableTransactionNumber(s *string) *UserPaymentPlanCreate {
	if s != nil {
		uppc.SetTransactionNumber(*s)
	}
	return uppc
}

// SetSourceAccountNumber sets the "source_account_number" field.
func (uppc *UserPaymentPlanCreate) SetSourceAccountNumber(s string) *UserPaymentPlanCreate {
	uppc.mutation.SetSourceAccountNumber(s)
	return uppc
}

// SetNillableSourceAccountNumber sets the "source_account_number" field if the given value is not nil.
func (uppc *UserPaymentPlanCreate) SetNillableSourceAccountNumber(s *string) *UserPaymentPlanCreate {
	if s != nil {
		uppc.SetSourceAccountNumber(*s)
	}
	return uppc
}

// SetDestinationAccountNumber sets the "destination_account_number" field.
func (uppc *UserPaymentPlanCreate) SetDestinationAccountNumber(s string) *UserPaymentPlanCreate {
	uppc.mutation.SetDestinationAccountNumber(s)
	return uppc
}

// SetNillableDestinationAccountNumber sets the "destination_account_number" field if the given value is not nil.
func (uppc *UserPaymentPlanCreate) SetNillableDestinationAccountNumber(s *string) *UserPaymentPlanCreate {
	if s != nil {
		uppc.SetDestinationAccountNumber(*s)
	}
	return uppc
}

// SetDeleted sets the "deleted" field.
func (uppc *UserPaymentPlanCreate) SetDeleted(b bool) *UserPaymentPlanCreate {
	uppc.mutation.SetDeleted(b)
	return uppc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (uppc *UserPaymentPlanCreate) SetNillableDeleted(b *bool) *UserPaymentPlanCreate {
	if b != nil {
		uppc.SetDeleted(*b)
	}
	return uppc
}

// SetCreatedAt sets the "created_at" field.
func (uppc *UserPaymentPlanCreate) SetCreatedAt(t time.Time) *UserPaymentPlanCreate {
	uppc.mutation.SetCreatedAt(t)
	return uppc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uppc *UserPaymentPlanCreate) SetNillableCreatedAt(t *time.Time) *UserPaymentPlanCreate {
	if t != nil {
		uppc.SetCreatedAt(*t)
	}
	return uppc
}

// SetID sets the "id" field.
func (uppc *UserPaymentPlanCreate) SetID(i int64) *UserPaymentPlanCreate {
	uppc.mutation.SetID(i)
	return uppc
}

// SetUserIDID sets the "user_id" edge to the User entity by ID.
func (uppc *UserPaymentPlanCreate) SetUserIDID(id int64) *UserPaymentPlanCreate {
	uppc.mutation.SetUserIDID(id)
	return uppc
}

// SetUserID sets the "user_id" edge to the User entity.
func (uppc *UserPaymentPlanCreate) SetUserID(u *User) *UserPaymentPlanCreate {
	return uppc.SetUserIDID(u.ID)
}

// SetPlanIDID sets the "plan_id" edge to the Plan entity by ID.
func (uppc *UserPaymentPlanCreate) SetPlanIDID(id int) *UserPaymentPlanCreate {
	uppc.mutation.SetPlanIDID(id)
	return uppc
}

// SetPlanID sets the "plan_id" edge to the Plan entity.
func (uppc *UserPaymentPlanCreate) SetPlanID(p *Plan) *UserPaymentPlanCreate {
	return uppc.SetPlanIDID(p.ID)
}

// Mutation returns the UserPaymentPlanMutation object of the builder.
func (uppc *UserPaymentPlanCreate) Mutation() *UserPaymentPlanMutation {
	return uppc.mutation
}

// Save creates the UserPaymentPlan in the database.
func (uppc *UserPaymentPlanCreate) Save(ctx context.Context) (*UserPaymentPlan, error) {
	uppc.defaults()
	return withHooks(ctx, uppc.sqlSave, uppc.mutation, uppc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uppc *UserPaymentPlanCreate) SaveX(ctx context.Context) *UserPaymentPlan {
	v, err := uppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uppc *UserPaymentPlanCreate) Exec(ctx context.Context) error {
	_, err := uppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uppc *UserPaymentPlanCreate) ExecX(ctx context.Context) {
	if err := uppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uppc *UserPaymentPlanCreate) defaults() {
	if _, ok := uppc.mutation.Deleted(); !ok {
		v := userpaymentplan.DefaultDeleted
		uppc.mutation.SetDeleted(v)
	}
	if _, ok := uppc.mutation.CreatedAt(); !ok {
		v := userpaymentplan.DefaultCreatedAt()
		uppc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uppc *UserPaymentPlanCreate) check() error {
	if _, ok := uppc.mutation.ReferenceNumber(); !ok {
		return &ValidationError{Name: "reference_number", err: errors.New(`ent: missing required field "UserPaymentPlan.reference_number"`)}
	}
	if _, ok := uppc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "UserPaymentPlan.deleted"`)}
	}
	if _, ok := uppc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserPaymentPlan.created_at"`)}
	}
	if _, ok := uppc.mutation.UserIDID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required edge "UserPaymentPlan.user_id"`)}
	}
	if _, ok := uppc.mutation.PlanIDID(); !ok {
		return &ValidationError{Name: "plan_id", err: errors.New(`ent: missing required edge "UserPaymentPlan.plan_id"`)}
	}
	return nil
}

func (uppc *UserPaymentPlanCreate) sqlSave(ctx context.Context) (*UserPaymentPlan, error) {
	if err := uppc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uppc.mutation.id = &_node.ID
	uppc.mutation.done = true
	return _node, nil
}

func (uppc *UserPaymentPlanCreate) createSpec() (*UserPaymentPlan, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPaymentPlan{config: uppc.config}
		_spec = sqlgraph.NewCreateSpec(userpaymentplan.Table, sqlgraph.NewFieldSpec(userpaymentplan.FieldID, field.TypeInt64))
	)
	if id, ok := uppc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uppc.mutation.Amount(); ok {
		_spec.SetField(userpaymentplan.FieldAmount, field.TypeInt64, value)
		_node.Amount = value
	}
	if value, ok := uppc.mutation.ReferenceNumber(); ok {
		_spec.SetField(userpaymentplan.FieldReferenceNumber, field.TypeString, value)
		_node.ReferenceNumber = value
	}
	if value, ok := uppc.mutation.TransactionNumber(); ok {
		_spec.SetField(userpaymentplan.FieldTransactionNumber, field.TypeString, value)
		_node.TransactionNumber = &value
	}
	if value, ok := uppc.mutation.SourceAccountNumber(); ok {
		_spec.SetField(userpaymentplan.FieldSourceAccountNumber, field.TypeString, value)
		_node.SourceAccountNumber = &value
	}
	if value, ok := uppc.mutation.DestinationAccountNumber(); ok {
		_spec.SetField(userpaymentplan.FieldDestinationAccountNumber, field.TypeString, value)
		_node.DestinationAccountNumber = &value
	}
	if value, ok := uppc.mutation.Deleted(); ok {
		_spec.SetField(userpaymentplan.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := uppc.mutation.CreatedAt(); ok {
		_spec.SetField(userpaymentplan.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := uppc.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userpaymentplan.UserIDTable,
			Columns: []string{userpaymentplan.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_userpaymentplans = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uppc.mutation.PlanIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userpaymentplan.PlanIDTable,
			Columns: []string{userpaymentplan.PlanIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.plan_userpaymentplans = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserPaymentPlanCreateBulk is the builder for creating many UserPaymentPlan entities in bulk.
type UserPaymentPlanCreateBulk struct {
	config
	err      error
	builders []*UserPaymentPlanCreate
}

// Save creates the UserPaymentPlan entities in the database.
func (uppcb *UserPaymentPlanCreateBulk) Save(ctx context.Context) ([]*UserPaymentPlan, error) {
	if uppcb.err != nil {
		return nil, uppcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uppcb.builders))
	nodes := make([]*UserPaymentPlan, len(uppcb.builders))
	mutators := make([]Mutator, len(uppcb.builders))
	for i := range uppcb.builders {
		func(i int, root context.Context) {
			builder := uppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPaymentPlanMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uppcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uppcb *UserPaymentPlanCreateBulk) SaveX(ctx context.Context) []*UserPaymentPlan {
	v, err := uppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uppcb *UserPaymentPlanCreateBulk) Exec(ctx context.Context) error {
	_, err := uppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uppcb *UserPaymentPlanCreateBulk) ExecX(ctx context.Context) {
	if err := uppcb.Exec(ctx); err != nil {
		panic(err)
	}
}
