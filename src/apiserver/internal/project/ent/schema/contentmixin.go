package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/google/uuid"
)

const (
	HeaderMaxLen = 1 << 5
)

type ContentMixin struct {
	mixin.Schema
}

func (ContentMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Immutable().
			Unique(),
		field.String("head").
			NotEmpty().
			MaxLen(HeaderMaxLen),
		field.Text("body").
			NotEmpty(),
	}
}
