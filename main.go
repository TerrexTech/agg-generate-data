package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/TerrexTech/agg-generate-data/collection"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/agg-generate-data/sample"
	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func validateEnv() error {
	missingVar, err := commonutil.ValidateEnv(
		"MONGO_HOSTS",
		"MONGO_DATABASE",
		"MONGO_INV_COLLECTION",
		"MONGO_DEVICE_COLLECTION",
		"MONGO_METRIC_COLLECTION",

		"MONGO_TIMEOUT",
	)

	if err != nil {
		err = errors.Wrapf(err, "Env-var %s is required for testing, but is not set", missingVar)
		return err
	}
	return nil
}

func main() {

	//Change sample size here
	// sampleSize := 3000

	log.Println("Reading environment file")
	err := godotenv.Load("./.env")
	if err != nil {
		err = errors.Wrap(err,
			".env file not found, env-vars will be read as set in environment",
		)
		log.Println(err)
	}

	err = validateEnv()
	if err != nil {
		log.Fatalln(err)
	}

	hosts := os.Getenv("MONGO_HOSTS")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("MONGO_DATABASE")
	collectionDev := os.Getenv("MONGO_DEVICE_COLLECTION")
	collectionInv := os.Getenv("MONGO_INV_COLLECTION")
	collectionMet := os.Getenv("MONGO_METRIC_COLLECTION")
	sampleSize := os.Getenv("SAMPLE_SIZE")

	// log.Println(hosts)
	config := mongo.ClientConfig{
		Hosts:               []string{hosts},
		Username:            username,
		Password:            password,
		TimeoutMilliseconds: 3000,
	}

	// ====> MongoDB Client
	client, err := mongo.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	invColl, err := collection.Inventory(client, database, collectionInv)
	if err != nil {
		log.Println(err)
	}

	deviceColl, err := collection.Device(client, database, collectionDev)
	if err != nil {
		log.Println(err)
	}

	metricColl, err := collection.Metric(client, database, collectionMet)
	if err != nil {
		log.Println(err)
	}

	//Generating and Inserting data
	inventory := []model.Inventory{}

	samplesize, err := strconv.Atoi(sampleSize)
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < samplesize; i++ {
		inv, err := sample.GetInventory()
		log.Println(inv.DateArrived)
		log.Println(inv.Timestamp, "Timestamp")
		log.Println(inv.ProjectedDate, "ProjectedDate")
		if err != nil {
			log.Println(err)
		}
		inventory = append(inventory, inv)
		time.Sleep(100 * time.Millisecond)
	}

	err = sample.InsertInventory(inventory, invColl)
	if err != nil {
		log.Println(err)
	}

	//For sold, donate, waste
	for _, v := range inventory {
		_, err := sample.GenSale(v, invColl)
		if err != nil {
			log.Println(err)
		}

		_, err = sample.GenDonate(v, invColl)
		if err != nil {
			log.Println(err)
		}
		_, err = sample.GenWaste(v, invColl)
		if err != nil {
			log.Println(err)
		}
		// _, err = sample.GenFlashSale(v, invColl)
		// if err != nil {
		// 	log.Println(err)
		// }
	}

	device := []model.Device{}
	for _, v := range inventory {
		dev, err := sample.GetDevice(v)
		if err != nil {
			log.Println(err)
		}
		device = append(device, dev)
	}

	err = sample.InsertDevice(device, deviceColl)
	if err != nil {
		log.Println(err)
	}

	metric := []model.Metric{}
	for _, v := range device {
		met, err := sample.GetMetric(v)
		if err != nil {
			log.Println(err)
		}
		metric = append(metric, met)
	}

	// log.Println(metric)

	err = sample.InsertMetric(metric, metricColl)
	if err != nil {
		log.Println(err)
	}

	// _, err = sample.GenDonate(v, invColl)
	// if err != nil {
	// 	log.Println(err)
	// }

	// fsale, err := sample.GetFlashSale(v)
	// if err != nil {
	// 	log.Println(err)
	// }

	// // inventory = append(inventory, sa)
	// inventory = append(inventory, fsale)

	// for _, v := range inventory {
	// 	log.Println(v)
	// }

	// for _, v := range device {
	// 	log.Println(v)
	// }

	// for _, v := range metric {
	// 	log.Println(v)
	// }

	// sale := []model.Sale{}
	// for _, v := range inventory {
	// 	sa, err := sample.GetSale(v)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	sale = append(sale, sa)
	// }

	// for _, v := range sale {
	// 	log.Println(v)
	// }

	// flashsale := []model.Inventory{}
	// for _, v := range inventory {
	// 	fsale, err := sample.GetFlashSale(v)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	flashsale = append(flashsale, fsale)
	// }

	// TestIfDataGenerated()

	// hosts := os.Getenv("MONGO_HOSTS")
	// username := os.Getenv("MONGO_USERNAME")
	// password := os.Getenv("MONGO_PASSWORD")
	// database := os.Getenv("MONGO_DATABASE")
	// // collectionFlash := os.Getenv("MONGO_FLASH_COLLECTION")
	// collectionInv := os.Getenv("MONGO_INV_COLLECTION")
	// // collectionMet := os.Getenv("MONGO_METRIC_COLLECTION")
	// collectionDev := os.Getenv("MONGO_DEVICE_COLLECTION")

	// // consumerEventgroup := os.Getenv("KAFKA_CONSUMER_EVENT_GROUP")
	// // consumerEventQueryGroup := os.Getenv("KAFKA_CONSUMER_EVENT_QUERY_GROUP")
	// // consumerEventTopic := os.Getenv("KAFKA_CONSUMER_EVENT_TOPIC")
	// // consumerEventQueryTopic := os.Getenv("KAFKA_CONSUMER_EVENT_QUERY_TOPIC")
	// // producerEventQueryTopic := os.Getenv("KAFKA_PRODUCER_EVENT_QUERY_TOPIC")
	// // producerResponseTopic := os.Getenv("KAFKA_PRODUCER_RESPONSE_TOPIC")

	// timeoutMilliStr := os.Getenv("MONGO_TIMEOUT")
	// parsedTimeoutMilli, err := strconv.Atoi(timeoutMilliStr)
	// if err != nil {
	// 	err = errors.Wrap(err, "Error converting Timeout value to int32")
	// 	log.Println(err)
	// 	log.Println("MONGO_TIMEOUT value will be set to 3000 as default value")
	// 	parsedTimeoutMilli = 3000
	// }
	// timeoutMilli := uint32(parsedTimeoutMilli)

	// log.Println(hosts)
	// // configFlash = DBIConfig{
	// // 	Hosts:               *commonutil.ParseHosts(hosts),
	// // 	Username:            username,
	// // 	Password:            password,
	// // 	TimeoutMilliseconds: timeoutMilli,
	// // 	Database:            database,
	// // 	Collection:          collectionFlash,
	// // }

	// // configMetric = DBIConfig{
	// // 	Hosts:               *commonutil.ParseHosts(hosts),
	// // 	Username:            username,
	// // 	Password:            password,
	// // 	TimeoutMilliseconds: timeoutMilli,
	// // 	Database:            database,
	// // 	Collection:          collectionMet,
	// // }

	// configInv := DBIConfig{
	// 	Hosts:               []string{hosts},
	// 	Username:            username,
	// 	Password:            password,
	// 	TimeoutMilliseconds: timeoutMilli,
	// 	Database:            database,
	// 	Collection:          collectionInv,
	// }

	// configDev := DBIConfig{
	// 	Hosts:               []string{hosts},
	// 	Username:            username,
	// 	Password:            password,
	// 	TimeoutMilliseconds: timeoutMilli,
	// 	Database:            database,
	// 	Collection:          collectionDev,
	// }

	// // dbFlash, err := GenerateDB(configFlash, &Flash{})
	// // if err != nil {
	// // 	err = errors.Wrap(err, "Error connecting to Inventory DB")
	// // 	log.Println(err)
	// // 	return
	// // }

	// // log.Println(configInv, configMetric)

	// // dbMetric, err := GenerateDB(configMetric, &Metric{})
	// // if err != nil {
	// // 	err = errors.Wrap(err, "Error connecting to Inventory DB")
	// // 	log.Println(err)
	// // 	return
	// // }

	// dbInventory, err := GenerateDB(configInv, &inventory.Inventory{})
	// if err != nil {
	// 	err = errors.Wrap(err, "Error connecting to Inventory DB")
	// 	log.Println(err)
	// 	return
	// }

	// // dbDevice, err := GenerateDB(configDev, &device.Device{})
	// // if err != nil {
	// // 	err = errors.Wrap(err, "Error connecting to Inventory DB")
	// // 	log.Println(err)
	// // 	return
	// // }

	// inventory := []inventory.Inventory{}
	// for i := 0; i < 100; i++ {
	// 	inventory = append(inventory, GenerateDataForInv())
	// }

	// insertResult, err := AddInventory(inventory, dbInventory)
	// if err != nil {
	// 	err = errors.Wrap(err, "Error inserting inventory in DB")
	// 	log.Println(err)
	// 	return
	// }

}
