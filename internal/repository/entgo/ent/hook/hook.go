// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/correlateme-server/internal/repository/entgo/ent"
)

// The CorrelationFunc type is an adapter to allow the use of ordinary
// function as Correlation mutator.
type CorrelationFunc func(context.Context, *ent.CorrelationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CorrelationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CorrelationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CorrelationMutation", m)
	}
	return f(ctx, mv)
}

// The DatasetFunc type is an adapter to allow the use of ordinary
// function as Dataset mutator.
type DatasetFunc func(context.Context, *ent.DatasetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DatasetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DatasetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DatasetMutation", m)
	}
	return f(ctx, mv)
}

// The IndicatorFunc type is an adapter to allow the use of ordinary
// function as Indicator mutator.
type IndicatorFunc func(context.Context, *ent.IndicatorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IndicatorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.IndicatorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IndicatorMutation", m)
	}
	return f(ctx, mv)
}

// The ObservationFunc type is an adapter to allow the use of ordinary
// function as Observation mutator.
type ObservationFunc func(context.Context, *ent.ObservationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ObservationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ObservationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ObservationMutation", m)
	}
	return f(ctx, mv)
}

// The ScaleFunc type is an adapter to allow the use of ordinary
// function as Scale mutator.
type ScaleFunc func(context.Context, *ent.ScaleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScaleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScaleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScaleMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
