package utils

import (
	"csv-to-mysql/table/defs"
	"database/sql"
	"fmt"
	"strings"
	"unicode/utf8"
)

func createTableStatement(modifiedHeaders []string) string {
	sqlTypes := defs.TableSQLTypes()
	impl := []string{}
	for i := 0; i < len(modifiedHeaders); i++ {
		impl = append(impl, modifiedHeaders[i]+" "+sqlTypes[modifiedHeaders[i]])
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS raw (%s)", strings.Join(impl, ","))
}

func quoteRespectingCommaSplit(line string) []string {
	linebytes := []byte(line)
	res := []string{}
	buffer := []rune{}
	isInQuotes := false

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

	// append remaining buffer, since character is not a ,
	res = append(res, string(buffer))

	return res
}

func ConvertCSVValuesToSQLValues(line string, quoteTypes []defs.TableQuoteType) string {

	fields := quoteRespectingCommaSplit(line)
	res := []string{}

	for i := range fields {
		field := fields[i]
		quoteType := quoteTypes[i]
		if quoteType == defs.InvalidHeader {
			res = append(res, "INVALID_HEADER_ERROR")
		} else if len(field) == 0 {
			res = append(res, "NULL")
		} else if (field[0] == '"' && quoteType == defs.Quoted) || quoteType == defs.Unquoted {
			// no modifications need to be made
			res = append(res, field)
		} else {
			res = append(res, "\""+field+"\"")
		}
	}

	return "(" + strings.Join(res, ",") + ")"
}

func InitializeDb(db *sql.DB, modifiedHeaders []string) {
	_, err := db.Exec("DROP TABLE IF EXISTS raw")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(createTableStatement(modifiedHeaders))
	if err != nil {
		panic(err)
	}
}

func insertIntoStatement(headers []string, values []string) string {
	query := fmt.Sprintf(
		"INSERT INTO raw\n(%s)\nVALUES\n%s\n",
		strings.Join(headers, ","),
		strings.Join(values, ","),
	)
	return query
}

func InsertInto(db *sql.DB, headers []string, values []string) {
	query := insertIntoStatement(headers, values)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
