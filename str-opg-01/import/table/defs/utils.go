package defs

func IntoTableQuoteType(headers []string) []TableQuoteType {
	types := TableQuoteTypesMap()
	res := []TableQuoteType{}
	for i := range headers {
		res = append(res, types[headers[i]])
	}
	return res
}
