package dblayer

import (
	"cloud-native-programming-with-golang/Chapter02/myevents/src/eventsservice/lib/persistence"
	"cloud-native-programming-with-golang/Chapter02/myevents/src/eventsservice/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
