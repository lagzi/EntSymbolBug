package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const (
	ComponentPictures = "pictures"
)

type Component struct {
	ent.Schema
}

func (Component) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Component) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("some_id").Unique(),
	}
}

func (Component) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(ComponentPictures, Picture.Type),
	}
}

func (Component) Indexes() []ent.Index {
	return []ent.Index{index.Fields("some_id")}
}
