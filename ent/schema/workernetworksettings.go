package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

const (
	edgeName = "worker_contained_information"
)

type WorkerNetworkSettings struct {
	ent.Schema
}

func (WorkerNetworkSettings) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges(edgeName),
	}
}

func (WorkerNetworkSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(edgeName, WorkerContainedInformation.Type).Ref(otherEdgeName).Unique(),
	}
}
