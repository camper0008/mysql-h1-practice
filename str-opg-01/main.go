package main

import (
	"fmt"
	"strings"
)

func splitLine(line string) []string {
	res := []string{}
	isInQuotes := false
	buffer := ""
	for i := 0; i < len(line); i++ {
		if line[i] == '"' {
			isInQuotes = !isInQuotes
		} else if line[i] == ',' && !isInQuotes {
			res = append(res, buffer)
			buffer = ""
		} else {
			buffer += string(line[i])
		}
	}
	return res
}

func convertCSVValuesToSQLValues(line string, quoteTypes []TableQuoteType) string {
	sep := splitLine(line)
	res := []string{}
	for i := range sep {
		v := sep[i]
		quoteType := quoteTypes[i]
		if quoteType == InvalidHeader {
			res = append(res, "INVALID_HEADER_ERROR")
		} else if len(v) == 0 {
			res = append(res, "NULL")
		} else if v[0] == '"' || quoteType == Unquoted {
			// no modifications need to be made
			res = append(res, v)
		} else {
			res = append(res, "\""+v+"\"")
		}
	}

	return "(" + strings.Join(res, ",") + ")"
}

func modifyHeader(filename string) []string {
	rawHeader := ReadCSVHeader(filename)

	splitHeader := strings.Split(rawHeader, ",")
	globalizedHeader := DanishToEnglishFieldNames(splitHeader)
	return globalizedHeader
}

func modifyValues(filename string, modifiedHeaders []string) []string {
	values := ReadCSVEverythingElse(filename)
	types := IntoTableQuoteType(modifiedHeaders)

	res := []string{}

	for i := range values {
		res = append(res, convertCSVValuesToSQLValues(values[i], types))
	}

	return res
}

func main() {
	const FILENAME string = "./address.csv"
	headers := modifyHeader(FILENAME)
	values := modifyValues(FILENAME, headers)

	fmt.Printf("INSERT INTO address\n(%s)\nVALUES\n", strings.Join(headers, ","))
	for i := 0; i < len(values) && i < 2; i++ {
		fmt.Println(values[i])
	}
}
