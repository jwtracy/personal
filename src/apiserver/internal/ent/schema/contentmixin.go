package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

const (
	HeaderMaxLen = 1 << 5
)

type ContentMixin struct {
	mixin.Schema
}

func (ContentMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("head").
			NotEmpty().
			MaxLen(HeaderMaxLen),
		field.Text("body").
			NotEmpty(),
	}
}
