package collection

import (
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

func Metric(client *mongo.Client, database string, collMet string) (*mongo.Collection, error) {
	conn := &mongo.ConnectionConfig{
		Client:  client,
		Timeout: 5000,
	}
	// Index Configuration
	indexConfigs := []mongo.IndexConfig{}

	// ====> Create New Collection
	c := &mongo.Collection{
		Connection:   conn,
		Name:         collMet,
		Database:     database,
		SchemaStruct: &model.Metric{},
		Indexes:      indexConfigs,
	}
	return mongo.EnsureCollection(c)
}
