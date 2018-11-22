package sample

import (
	"github.com/TerrexTech/agg-generate-data/mockutil"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	mgo "github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

func GenWaste(inv model.Inventory, invColl *mongo.Collection) (*mgo.UpdateResult, error) {
	var wasteWeight float64

	totalWeight := inv.TotalWeight
	cumulativeWeight := inv.DonateWeight + inv.SoldWeight + inv.FlashSaleWeight + inv.WasteWeight

	// if inv.WasteWeight < inv.TotalWeight && inv.WasteWeight == 0 && inv.DonateWeight < (inv.TotalWeight+inv.SoldWeight) {
	// 	wasteWeight = mockutil.GenFloat(inv.WasteWeight+1, 20)
	// }

	// if inv.WasteWeight < inv.TotalWeight && inv.WasteWeight < (inv.TotalWeight+inv.SoldWeight) && inv.WasteWeight != 0 {
	// 	wasteWeight = mockutil.GenFloat(inv.WasteWeight, 30)
	// }

	if cumulativeWeight < totalWeight {
		wasteWeight = mockutil.GenFloat(inv.WasteWeight+1, 30)
	}

	itemID := inv.ItemID.String()

	filter := map[string]string{
		"itemID": itemID,
	}

	update := &map[string]interface{}{
		"wasteWeight": wasteWeight,
	}
	updateResult, err := invColl.UpdateMany(filter, update)
	if err != nil {
		err = errors.Wrap(err, "Unable to update event - Sale")
		return nil, err
	}

	return updateResult, nil
}
