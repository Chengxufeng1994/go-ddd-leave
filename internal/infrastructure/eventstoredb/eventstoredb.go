package eventstoredb

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
)

func NewEventStoreDB() (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString("esdb://192.168.1.230:2113?tls=false")
	if err != nil {
		return nil, err
	}
	return esdb.NewClient(settings)
}
