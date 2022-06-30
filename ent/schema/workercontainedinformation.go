package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

const (
	otherEdgeName = "network_settings"
)

type WorkerContainedInformation struct {
	ent.Schema
}

func (WorkerContainedInformation) Edges() []ent.Edge {
	edges := []ent.Edge{
		edge.To(otherEdgeName, WorkerNetworkSettings.Type).Unique(),
	}

	return edges
}
