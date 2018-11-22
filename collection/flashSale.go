package collection

// import (
// 	"github.com/TerrexTech/agg-generate-data/model"
// 	"github.com/TerrexTech/go-mongoutils/mongo"
// )

// func FlashSale(client *mongo.Client) (*mongo.Collection, error) {
// 	conn := &mongo.ConnectionConfig{
// 		Client:  client,
// 		Timeout: 5000,
// 	}
// 	// Index Configuration
// 	indexConfigs := []mongo.IndexConfig{}

// 	// ====> Create New Collection
// 	c := &mongo.Collection{
// 		Connection:   conn,
// 		Name:         "agg_flashsale",
// 		Database:     "rns_projections",
// 		SchemaStruct: &model.Inventory{},
// 		Indexes:      indexConfigs,
// 	}
// 	return mongo.EnsureCollection(c)
// }
