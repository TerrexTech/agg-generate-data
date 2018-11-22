package sample

import (
	"log"

	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/pkg/errors"
)

func GetDevice(inv model.Inventory) (model.Device, error) {

	return model.Device{
		ItemID:        inv.ItemID,
		DeviceID:      inv.DeviceID,
		DateInstalled: inv.Timestamp,
		Lot:           inv.Lot,
		Name:          inv.Name,
		Status:        "Healthy",
		SKU:           inv.SKU,
	}, nil
}

func InsertDevice(dev []model.Device, devColl *mongo.Collection) error {
	for _, v := range dev {
		_, err := devColl.InsertOne(v)
		if err != nil {
			err = errors.Wrap(err, "Unable to insert data - device")
			log.Println(err)
			return err
		}
		// log.Println(insertResult)
	}
	return nil
}
