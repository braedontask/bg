package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BioUser holds the schema definition for the BioUser entity.
type BioUser struct {
	ent.Schema
}

// Fields of the BioUser.
func (BioUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the BioUser.
func (BioUser) Edges() []ent.Edge {
	return nil
}
