package sample

import (
	"log"
	"time"

	"github.com/TerrexTech/agg-generate-data/mockutil"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/pkg/errors"
)

func GetMetric(device model.Device) (model.Metric, error) {
	co2 := mockutil.GenFloat(400, 1200)

	metricID := mockutil.GenUUID().String()
	itemID := device.ItemID.String()
	deviceID := device.DeviceID.String()

	return model.Metric{
		MetricID:      metricID,
		ItemID:        itemID,
		DeviceID:      deviceID,
		Timestamp:     time.Now().Unix(),
		TempIn:        mockutil.GenFloat(21, 27),
		Humidity:      mockutil.GenFloat(40, 80),
		CarbonDioxide: co2,
		Ethylene:      co2 / 3,
		SKU:           device.SKU,
		Name:          device.Name,
		Lot:           device.Lot,
	}, nil
}

func InsertMetric(met []model.Metric, metColl *mongo.Collection) error {
	for _, v := range met {
		log.Println(v)
		insertResult, err := metColl.InsertOne(v)
		if err != nil {
			err = errors.Wrap(err, "Unable to insert data - metric")
			log.Println(err)
			return err
		}
		log.Println(insertResult)
	}
	return nil
}
