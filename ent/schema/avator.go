package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Avator holds the schema definition for the Avator entity.
type Avator struct {
	ent.Schema
}

// Fields of the Avator.
func (Avator) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("bio"),
	}
}

// Edges of the Avator.
func (Avator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("avators").
			Unique(),
	}
}
