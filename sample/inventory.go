package sample

import (
	"log"
	"time"

	"github.com/TerrexTech/agg-generate-data/mockutil"
	"github.com/TerrexTech/agg-generate-data/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/pkg/errors"
)

func GetInventory() (model.Inventory, error) {
	barcode := mockutil.GenBarcode()
	name := mockutil.GenFruitName()
	dateArr := mockutil.GenTime()
	timestamp := dateArr + int64(mockutil.GenInt(7200, 15000))
	projectedDate := time.Unix(dateArr, 0).AddDate(0, 0, mockutil.GenInt(7, 15)).Unix()

	return model.Inventory{
		ItemID: mockutil.GenUUID(),
		// Barcode:      mockutil.GenBarcode(),
		DateArrived:   dateArr,
		DeviceID:      mockutil.GenUUID(),
		Lot:           mockutil.GenLot(),
		Name:          name,
		Origin:        mockutil.GenOrigin(),
		Price:         mockutil.GenFloat(1, 10),
		RSCustomerID:  mockutil.GenUUID(),
		SKU:           mockutil.SkuDict(name),
		Timestamp:     timestamp,
		TotalWeight:   mockutil.GenFloat(300, 1000),
		UPC:           barcode,
		ProjectedDate: projectedDate,
	}, nil
}

func InsertInventory(inv []model.Inventory, invColl *mongo.Collection) error {
	for _, v := range inv {
		_, err := invColl.InsertOne(v)
		if err != nil {
			err = errors.Wrap(err, "Unable to insert data - inv")
			log.Println(err)
			return err
		}
		// log.Println(insertResult)
	}
	return nil
}

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// // type ModifyInvData struct {
// // 	Inv        Inventory
// // 	Datearr    int64
// // 	Expirydate int64
// // 	Timestamp  int64
// // 	Randnum    int64
// // }

// func random(min, max int64) int64 {
// 	return rand.Int63n(max-min) + min
// }

// func generateRandomValue(num1, num2 int64) int64 {
// 	// rand.Seed(time.Now().Unix())
// 	return random(num1, num2)
// }

// func generateNewUUID() uuuid.UUID {
// 	uuid, err := uuuid.NewV4()
// 	if err != nil {
// 		err = errors.Wrap(err, "Unable to generate UUID")
// 		log.Println(err)
// 	}
// 	return uuid
// }

// func GenFloat(min float64, max float64) float64 {
// 	s1 := rand.NewSource(time.Now().UnixNano())
// 	r1 := rand.New(s1)
// 	random := min + r1.Float64()*(max-min)
// 	return random
// }

// func GenString(chars string, length int) string {
// 	if chars == "" {
// 		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
// 	}
// 	charLength := len(chars)

// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)

// 	// Note that err == nil only if we read len(b) bytes.
// 	if err != nil {
// 		err = errors.Wrap(err, "Error while generating random String")
// 		log.Println(err)
// 		return ""
// 	}
// 	for i, b := range bytes {
// 		bytes[i] = chars[b%byte(charLength)]
// 	}

// 	return string(bytes)
// }

// func GenSKU() string {
// 	s1 := GenString("", 3)
// 	s2 := GenString("", 3)
// 	s3 := GenString("", 3)
// 	s4 := GenString("", 2)

// 	return fmt.Sprintf("%s-%s-%s-%s", s1, s2, s3, s4)
// }

// var productsName = []string{"Banana", "Orange", "Apple", "Mango", "Strawberry", "Tomato", "Lettuce", "Pear", "Grapes", "Sweet Pepper"}
// var lot = []string{"A101", "B201", "O301", "M401", "S501", "T601", "L701", "P801", "G901", "SW1001"}
// var provinceNames = []string{"ON Canada", "BC Canada", "SK Canada", "MN Canada", "NS Canada", "PEI Canada", "QC Canada"}

// func GenerateDataForInv() inventory.Inventory {

// 	randNameAndLocation := generateRandomValue(1, 10)
// 	randOrigin := generateRandomValue(1, 6)
// 	randDateArr := generateRandomValue(1, 6)                          //in hours
// 	randTimestamp := generateRandomValue(randDateArr, randDateArr+1)  //in hours
// 	randExpiry := generateRandomValue(((randTimestamp / 24) + 1), 21) //in days
// 	randDatesold := generateRandomValue(randTimestamp, randExpiry*24) //in hours
// 	randPrice := GenFloat(0.5, 9)
// 	randTotalWeight := generateRandomValue(100, 300)
// 	randWasteWeight := generateRandomValue(1, 80)
// 	// randProdQuan := generateRandomValue(10, 200)
// 	itemId := generateNewUUID()
// 	sku := GenSKU()
// 	deviceId := generateNewUUID()
// 	lot := lot[randNameAndLocation]
// 	customerId := generateNewUUID()
// 	timestamp := time.Now().Unix()

// 	inventory := inventory.Inventory{
// 		ItemID:       itemId,
// 		UPC:          sku,
// 		SKU:          sku,
// 		RSCustomerID: customerId,
// 		DeviceID:     deviceId,
// 		Name:         productsName[randNameAndLocation-1], //-1 because rand starts from 1
// 		Origin:       provinceNames[randOrigin-1],
// 		TotalWeight:  float64(randTotalWeight),
// 		Price:        float64(randPrice),
// 		Lot:          lot,
// 		WasteWeight:  float64(randWasteWeight - 1),
// 		DonateWeight: float64(generateRandomValue(1, 21)),
// 		DateArrived:  time.Now().Add(time.Duration(randDateArr) * time.Hour).Unix(),
// 		// ExpiryDate:   time.Now().AddDate(0, 0, int(randExpiry)).Unix(),
// 		// Timestamp:    time.Now().Add(time.Duration(randTimestamp) * time.Hour).Unix(),
// 		Timestamp: timestamp,
// 		// DateSold:     time.Now().Add(time.Duration(randDatesold) * time.Hour).Unix(),
// 		DateSold: time.Now().Add(time.Duration(randDatesold) * time.Hour).Unix(),

// 		SalePrice:  float64(generateRandomValue(2, 4)),
// 		SoldWeight: float64(generateRandomValue(randWasteWeight, randTotalWeight)),
// 	}
// 	return inventory
// }

// func GenerateDataForDevice() {

// }

// func GenFakeBarcode(barType string) string {
// 	var num int64
// 	var t string
// 	if barType == "upc" {
// 		num = generateRandomValue(111111111111, 999999999999)
// 		t = strconv.Itoa(int(num))
// 	}
// 	if barType == "sku" {
// 		num = generateRandomValue(11111111, 99999999)
// 		t = strconv.Itoa(int(num))
// 	}
// 	return t
// }

// func TestIfDataGenerated() {
// 	inventory := []inventory.Inventory{}
// 	for i := 0; i < 100; i++ {
// 		inventory = append(inventory, GenerateDataForInv())
// 	}

// 	log.Println(inventory)

// 	// jsonWithInvData, err := json.Marshal(&inventory)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	// log.Println(jsonWithInvData)
// }
