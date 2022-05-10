package main

import (
	"csv-to-mysql/db_api"
	db_utils "csv-to-mysql/db_api/utils"
	"csv-to-mysql/fs"
	"csv-to-mysql/table/defs"
	"csv-to-mysql/utils"
	"fmt"
	"strings"
)

func ModifyHeader(filename string) []string {
	rawHeader := fs.ReadCSVHeader(filename)

	splitHeader := strings.Split(rawHeader, ",")
	globalizedHeader := utils.DanishToEnglishFieldNames(splitHeader)
	return globalizedHeader
}

func ModifyValues(filename string, modifiedHeaders []string) []string {
	values := fs.ReadCSVEverythingElse(filename)
	types := defs.IntoTableQuoteType(modifiedHeaders)

	res := []string{}

	for i := range values {
		res = append(res, db_utils.ConvertCSVValuesToSQLValues(values[i], types))
	}

	return res
}

func main() {
	const FILENAME string = "./address.csv"
	const VALUES_PER_ITERATION int = 1000

	headers := ModifyHeader(FILENAME)
	values := ModifyValues(FILENAME, headers)

	fmt.Println("connecting to db...")
	db := db_api.Connect()
	fmt.Println("done!")
	fmt.Println("initializing db...")
	db_utils.InitializeDb(db, headers)
	fmt.Println("done!")

	fmt.Println("inserting values...")
	for i := 0; i < len(values); i += VALUES_PER_ITERATION {
		lowest := utils.Min(i+VALUES_PER_ITERATION, len(values))

		db_utils.InsertInto(db, headers, values[i:lowest])

		percentage := float64(i) / float64(len(values)) * 100

		fmt.Printf("%0.2f%% [%d / %d]\n", percentage, i, len(values))
	}

	fmt.Printf("%0.2f%% [%d / %d]\n", 100.0, len(values), len(values))
	fmt.Println("done!")

	fmt.Println("inserting post code map table...")
	db_utils.PostCodeMap(db)
	fmt.Println("done!")

	fmt.Println("inserting region code map table...")
	db_utils.RegionCodeMap(db)
	fmt.Println("done!")

	fmt.Println("inserting address table...")
	db_utils.Address(db)
	fmt.Println("done!")

	fmt.Println("inserting postcard view...")
	db_utils.Postcard(db)
	fmt.Println("done!")

	db.Close()
}
