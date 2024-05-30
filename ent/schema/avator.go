package schema

import "entgo.io/ent"

// Avator holds the schema definition for the Avator entity.
type Avator struct {
	ent.Schema
}

// Fields of the Avator.
func (Avator) Fields() []ent.Field {
	return nil
}

// Edges of the Avator.
func (Avator) Edges() []ent.Edge {
	return nil
}
