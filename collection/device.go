package collection

import (
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

func Device(client *mongo.Client, database string, collDev string) (*mongo.Collection, error) {
	conn := &mongo.ConnectionConfig{
		Client:  client,
		Timeout: 5000,
	}
	// Index Configuration
	indexConfigs := []mongo.IndexConfig{}

	// ====> Create New Collection
	c := &mongo.Collection{
		Connection:   conn,
		Name:         collDev,
		Database:     database,
		SchemaStruct: &model.Device{},
		Indexes:      indexConfigs,
	}
	return mongo.EnsureCollection(c)
}
