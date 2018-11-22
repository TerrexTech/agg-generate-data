package collection

import (
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

func Inventory(client *mongo.Client, database string, collInv string) (*mongo.Collection, error) {
	conn := &mongo.ConnectionConfig{
		Client:  client,
		Timeout: 5000,
	}
	// Index Configuration
	indexConfigs := []mongo.IndexConfig{}

	// ====> Create New Collection
	c := &mongo.Collection{
		Connection:   conn,
		Name:         collInv,
		Database:     database,
		SchemaStruct: &model.Inventory{},
		Indexes:      indexConfigs,
	}
	return mongo.EnsureCollection(c)
}
