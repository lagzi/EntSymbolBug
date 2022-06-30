// Code generated by ent, DO NOT EDIT.

package workernetworksettings

const (
	// Label holds the string label denoting the workernetworksettings type in the database.
	Label = "worker_network_settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeWorkerContainedInformation holds the string denoting the worker_contained_information edge name in mutations.
	EdgeWorkerContainedInformation = "worker_contained_information"
	// Table holds the table name of the workernetworksettings in the database.
	Table = "worker_network_settings"
	// WorkerContainedInformationTable is the table that holds the worker_contained_information relation/edge.
	WorkerContainedInformationTable = "worker_network_settings"
	// WorkerContainedInformationInverseTable is the table name for the WorkerContainedInformation entity.
	// It exists in this package in order to avoid circular dependency with the "workercontainedinformation" package.
	WorkerContainedInformationInverseTable = "worker_contained_informations"
	// WorkerContainedInformationColumn is the table column denoting the worker_contained_information relation/edge.
	WorkerContainedInformationColumn = "worker_contained_information_network_settings"
)

// Columns holds all SQL columns for workernetworksettings fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "worker_network_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"worker_contained_information_network_settings",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}