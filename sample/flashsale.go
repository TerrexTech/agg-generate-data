package sample

import (
	"time"

	"github.com/TerrexTech/agg-generate-data/mockutil"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	mgo "github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

func GenFlashSale(inv model.Inventory, invColl *mongo.Collection) (*mgo.UpdateResult, error) {
	var flashSaleWeight float64

	// flashSaleWeight = mockutil.GenFloat(1, 30)

	cumulativeWeight := flashSaleWeight + inv.SoldWeight + inv.DonateWeight + inv.WasteWeight

	if cumulativeWeight > inv.TotalWeight {
		err := errors.New("Unable to update event - FlashSale")
		return nil, err
	}

	if cumulativeWeight < inv.TotalWeight && inv.TotalWeight-cumulativeWeight > 30 {
		flashSaleWeight = mockutil.GenFloat(1, 30)
	}

	itemID := inv.ItemID.String()

	filter := map[string]string{
		"itemID": itemID,
	}

	update := &map[string]interface{}{
		"flashSaleWeight":    flashSaleWeight,
		"flashSaleTimestamp": time.Now().Unix(),
	}
	updateResult, err := invColl.UpdateMany(filter, update)
	if err != nil {
		err = errors.Wrap(err, "Unable to update event - Sale")
		return nil, err
	}

	return updateResult, nil
}

// func GetFlashSale(inv *model.Inventory) (*model.Inventory, error) {
// 	var flashSaleWeight float64
// 	var flashTimestamp int64

// 	randNum := mockutil.GenInt(1, 4)
// 	totalSeconds := randNum * 3600

// 	flashTimestamp = ((time.Now().Unix()) - int64(totalSeconds))

// 	if inv.SoldWeight > 0 && inv.SoldWeight < inv.TotalWeight {
// 		flashSaleWeight = mockutil.GenFloat(inv.SoldWeight, inv.TotalWeight)
// 	}
// 	return &model.Inventory{
// 		FlashSaleWeight:    flashSaleWeight,
// 		FlashSaleTimestamp: flashTimestamp,
// 		OnFlashSale:        false,
// 	}, nil
// }
