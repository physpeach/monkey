package object

import (
	"fmt"
)

type ObjectType string

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	BOOLEAN_OBJ      = "BOOLEAN"
	INTEGER_OBJ      = "INTEGER"
	NULL_OBJ         = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
