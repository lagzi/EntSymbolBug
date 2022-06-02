package schema

import (
	"entgo.io/ent"
	edge "entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

const (
	UserToPictureEdge = "user_to_picture_edge_very_long_name_longer_than_the_amount_that_postgres_can_really_handle"
	//UserToPictureEdge = "short_edge"
	pictureEdge = "picture"
)

// User holds the schema definition for the User entity.
type Picture struct {
	ent.Schema
}

// Fields of the User.
func (Picture) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
	}
}

//func (Picture) Indexes() []ent.Index {
//	return []ent.Index{
//		index.Edges(pictureEdge),
//		index.Fields("key").
//			Edges(pictureEdge).
//			Unique(),
//	}
//}

// Edges of the User.
func (Picture) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(pictureEdge, User.Type).Ref(UserToPictureEdge),
	}
}
