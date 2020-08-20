package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/mixin"
)

// BlogPost holds the schema definition for the BlogPost entity.
type BlogPost struct {
	ent.Schema
}

func (BlogPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags", Topic.Type).
			Ref("blog_posts"),
	}
}

func (BlogPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ContentMixin{},
		mixin.Time{},
	}
}
