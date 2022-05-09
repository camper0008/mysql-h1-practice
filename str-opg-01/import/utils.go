package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func splitLine(line string) []string {
	linebytes := []byte(line)
	res := []string{}
	isInQuotes := false
	buffer := []rune{}

	for len(linebytes) > 0 {
		r, size := utf8.DecodeRune(linebytes)

		if r == '"' {
			isInQuotes = !isInQuotes
		} else if r == ',' && !isInQuotes {
			res = append(res, string(buffer))
			buffer = []rune{}
		} else {
			buffer = append(buffer, r)
		}

		linebytes = linebytes[size:]
	}

	/*
		for i := 0; i < len(line); i++ {
			if line[i] == '"' {
				isInQuotes = !isInQuotes
			} else if line[i] == ',' && !isInQuotes {
				res = append(res, string(buffer))
				buffer = []byte{}
			} else {
				r, _ := utf8.DecodeRune(linebyte)
				buffer = append(buffer, r)
			}
		}
	*/
	res = append(res, string(buffer))
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

func ModifyHeader(filename string) []string {
	rawHeader := ReadCSVHeader(filename)

	splitHeader := strings.Split(rawHeader, ",")
	globalizedHeader := DanishToEnglishFieldNames(splitHeader)
	return globalizedHeader
}

func ModifyValues(filename string, modifiedHeaders []string) []string {
	values := ReadCSVEverythingElse(filename)
	types := IntoTableQuoteType(modifiedHeaders)

	res := []string{}

	for i := range values {
		res = append(res, convertCSVValuesToSQLValues(values[i], types))
	}

	return res
}

func CreateStatement(modifiedHeaders []string) string {
	sqlTypes := TableSQLTypes()
	impl := []string{}
	for i := 0; i < len(modifiedHeaders); i++ {
		impl = append(impl, modifiedHeaders[i]+" "+sqlTypes[modifiedHeaders[i]])
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS address (%s)", strings.Join(impl, ","))
}
