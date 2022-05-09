package main

import (
	"fmt"
	"strings"
)

func main() {
	const FILENAME string = "./address.csv"
	headers := ModifyHeader(FILENAME)
	values := ModifyValues(FILENAME, headers)

	db := DbConnect()

	_, err := db.Exec("DROP TABLE IF EXISTS address")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(CreateStatement(headers))
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(values); i += 100 {
		lowest := i + 100
		if lowest > len(values) {
			lowest = len(values)
		}
		query := fmt.Sprintf(
			"INSERT INTO address\n(%s)\nVALUES\n%s\n",
			strings.Join(headers, ","),
			strings.Join(values[i:lowest], ",\n"),
		)

		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}

	defer db.Close()
}
