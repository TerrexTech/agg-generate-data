package mockutil

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/TerrexTech/uuuid"
	"github.com/pkg/errors"
)

func GenUUID() uuuid.UUID {
	uuid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "Error generating UUID")
		return uuuid.UUID{}
	}
	return uuid
}

func GenInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func GenFloat(min float64, max float64) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	// time.Sleep(1 * time.Millisecond)
	r1 := rand.New(s1)
	random := min + r1.Float64()*(max-min)
	return random
}

func GenString(chars string, length int) string {
	if chars == "" {
		chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	charLength := len(chars)

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)

	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		err = errors.Wrap(err, "Error while generating random String")
		log.Println(err)
		return ""
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(charLength)]
	}

	return string(bytes)
}

func GenBarcode() string {
	chars := "0123456789"
	barcode := GenString(chars, 8)
	return barcode
}

func GenSKU() string {
	s1 := GenString("", 2)
	s2 := GenString("", 2)
	s3 := GenString("", 2)
	s4 := GenString("", 1)

	return fmt.Sprintf("%s%s%s%s", s1, s2, s3, s4)
}

func SkuDict(item string) string {
	dictionary := map[string]string{
		"Apple":        "WLRPYUD",
		"Banana":       "KTYRZQD",
		"Grapes":       "UZTFEHB",
		"Lettuce":      "SQVZDMI",
		"Mango":        "ADNEBZE",
		"Orange":       "AWRFVYS",
		"Pear":         "COIQCIE",
		"Strawberry":   "HPFUXBX",
		"Sweet Pepper": "BBYITCE",
		"Tomato":       "PFJQUIN",
	}
	return dictionary[item]
}

func AvgWeight(fruit string) float64 {
	dictionary := map[string]float64{
		"Apple":        0.33,
		"Banana":       0.26,
		"Grapes":       0.01,
		"Lettuce":      0.028,
		"Mango":        0.44,
		"Orange":       0.30,
		"Pear":         0.44,
		"Strawberry":   0.026,
		"Sweet Pepper": 0.992,
		"Tomato":       0.328,
	}

	if dictionary[fruit] == 0 {
		return GenFloat(0, 1)
	}
	return dictionary[fruit]
}

func GenFruitName() string {
	dictionary := []string{
		"Apple",
		"Banana",
		"Grapes",
		"Lettuce",
		"Mango",
		"Orange",
		"Pear",
		"Strawberry",
		"Sweet Pepper",
		"Tomato",
	}

	index := GenInt(0, len(dictionary))
	return dictionary[index]
}

func GenOrigin() string {
	dictionary := []string{
		"ON Canada",
		"BC Canada",
		"SK Canada",
		"MN Canada",
		"NS Canada",
		"PEI Canada",
		"QC Canada",
	}

	index := GenInt(0, len(dictionary))
	return dictionary[index]
}

func GenLot() string {
	s1 := GenString("", 2)
	i1 := GenInt(0, 19)

	return fmt.Sprintf("%s%d", s1, i1)

	// dictionary := map[string]string{
	// 	"Apple":        "KB55",
	// 	"Banana":       "AH62",
	// 	"Grapes":       "8W74",
	// 	"Lettuce":      "4H74",
	// 	"Mango":        "W274",
	// 	"Orange":       "XP74",
	// 	"Pear":         "K174",
	// 	"Strawberry":   "4G74",
	// 	"Sweet Pepper": "F544",
	// 	"Tomato":       "TF44",
	// 	"Apple":        "TF44",
	// 	"Banana":       "F544",
	// 	"Grapes":       "4G74",
	// 	"Lettuce":      "K174",
	// 	"Mango":        "XP74",
	// 	"Orange":       "W274",
	// 	"Pear":         "4H74",
	// 	"Strawberry":   "8W74",
	// 	"Sweet Pepper": "AH62",
	// 	"Tomato":       "KB55",
	// }
	// return dictionary[item]

}

func GenTime() int64 {
	timeGen := time.Now().AddDate(0, 0, GenInt(-6, -1)).Unix()
	return timeGen
}
