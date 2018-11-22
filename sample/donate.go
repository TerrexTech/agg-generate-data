package sample

import (
	"github.com/TerrexTech/agg-generate-data/mockutil"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	mgo "github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

func GenDonate(inv model.Inventory, invColl *mongo.Collection) (*mgo.UpdateResult, error) {
	var donateWeight float64

	totalWeight := inv.TotalWeight
	cumulativeWeight := inv.DonateWeight + inv.SoldWeight + inv.FlashSaleWeight + inv.WasteWeight

	// if inv.DonateWeight < inv.TotalWeight && inv.DonateWeight == 0 && inv.DonateWeight < (inv.TotalWeight+inv.SoldWeight) {
	// 	donateWeight = mockutil.GenFloat(inv.DonateWeight+1, 30)
	// }

	// if inv.DonateWeight < inv.TotalWeight && inv.DonateWeight < (inv.TotalWeight+inv.SoldWeight) && inv.SoldWeight != 0 {
	// 	donateWeight = mockutil.GenFloat(inv.DonateWeight, 30)
	// }

	if cumulativeWeight < totalWeight {
		donateWeight = mockutil.GenFloat(inv.DonateWeight+1, 30)
	}

	itemID := inv.ItemID.String()

	filter := map[string]string{
		"itemID": itemID,
	}

	update := &map[string]interface{}{
		"donateWeight": donateWeight,
	}
	updateResult, err := invColl.UpdateMany(filter, update)
	if err != nil {
		err = errors.Wrap(err, "Unable to update event - Sale")
		return nil, err
	}

	return updateResult, nil
}
