package model

const (
	DATASTORE_OPERATION_INSERT = "INSERT"
	DATASTORE_OPERATION_UPDATE = "UPDATE"
	DATASTORE_OPERATION_DELETE = "DELETE"
	DATASTORE_OPERATION_SELECT = "SELECT"
)

type DatastoreSegement struct {
	Collection string
	Operation  string
	Query      string
	Parameter  map[string]interface{}
}
