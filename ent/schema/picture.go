package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const (
	userEdge       = "user"
	componentEdge  = "component"
	OtherField     = "other_important_thing"
	timestampField = "timestamp"
)

type Picture struct {
	ent.Schema
}

func (Picture) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges(componentEdge, userEdge).Fields(OtherField).Unique(),
		index.Fields(OtherField),
		index.Fields(OtherField, timestampField),
	}
}

func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.String(OtherField),
		field.Time(timestampField),
	}
}

func (Picture) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
func (Picture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(userEdge, User.Type).Ref(UserPicture).Unique(),
		edge.From(componentEdge, Component.Type).Ref(ComponentPictures).Unique(),
	}
}
