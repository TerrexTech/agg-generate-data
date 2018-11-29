package sample

import (
	"time"

	"github.com/TerrexTech/agg-generate-data/mockutil"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	mgo "github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

func GenSale(inv model.Inventory, invColl *mongo.Collection) (*mgo.UpdateResult, error) {
	var soldWeight float64

	soldWeight = mockutil.GenFloat(inv.SoldWeight, inv.TotalWeight)
	flashWeight := mockutil.GenFloat(1, soldWeight)

	itemID := inv.ItemID.String()

	filter := map[string]string{
		"itemID": itemID,
	}

	update := &map[string]interface{}{
		"soldWeight":      soldWeight,
		"flashSaleWeight": flashWeight,
		"dateSold":        time.Now().Unix(),
	}
	updateResult, err := invColl.UpdateMany(filter, update)
	if err != nil {
		err = errors.Wrap(err, "Unable to update event - Sale")
		return nil, err
	}

	return updateResult, nil

	// items := []model.SoldItem{}

	// numItems := mockutil.GenInt(1, 10)
	// for i := 0; i < numItems; i++ {
	// 	inv, err := mockutil.RandomInventory(invColl)
	// 	if err != nil {
	// 		err = errors.Wrap(err, "Error generating random inventory")
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	soldItem := model.SoldItem{
	// 		ItemID: inv.ItemID,
	// 		UPC:    inv.UPC,
	// 		Weight: mockutil.GenFloat(1, inv.TotalWeight-1),
	// 		Lot:    inv.Lot,
	// 		SKU:    inv.SKU,
	// 	}
	// 	items = append(items, soldItem)
	// }

	// soldItem := model.SoldItem{
	// 	ItemID: inv.ItemID,
	// 	UPC:    inv.UPC,
	// 	Weight: mockutil.GenFloat(1, inv.TotalWeight-1),
	// 	Lot:    inv.Lot,
	// 	SKU:    inv.SKU,
	// }
	// items = append(items, soldItem)

	// return model.Sale{
	// 	SaleID:    mockutil.GenUUID(),
	// 	Items:     items,
	// 	Timestamp: time.Now().Unix(),
	// }, nil
}

// func InsertSale(inv []model.Inventory, invColl *mongo.Collection) error {

// 	params := map[string]interface{}{
// 		"itemID": map[string]interface{}{
// 			"$eq": inv.ItemID,
// 		},
// 	}

// 	results, err := invColl.Find(params)
// 	if err != nil {
// 		err = errors.Wrap(err, "RandomDevice: Error in Find")
// 		log.Println(err)
// 		return nil, err
// 	}

// 	for _, v := range inv {
// 		insertResult, err := invColl.InsertOne(v)
// 		if err != nil {
// 			err = errors.Wrap(err, "Unable to insert data in inventory for soldWeight")
// 			log.Println(err)
// 			return err
// 		}
// 		log.Println(insertResult)
// 	}
// 	return nil
// }
