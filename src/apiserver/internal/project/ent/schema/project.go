package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Time("started"),
		field.Time("completed").
			Optional(),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags", Topic.Type).
			Ref("projects"),
	}
}

// Mixins of the Project.
func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ContentMixin{},
		mixin.Time{},
	}
}
