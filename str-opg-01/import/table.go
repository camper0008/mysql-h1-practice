package main

import (
	"strings"
)

type TableQuoteType int

const (
	InvalidHeader TableQuoteType = iota
	Quoted
	Unquoted
)

func DanishToEnglishFieldNames(names []string) []string {
	res := []string{}
	for i := 0; i < len(names); i++ {
		newStr := names[i]
		newStr = strings.Replace(newStr, "æ", "ae", -1)
		newStr = strings.Replace(newStr, "ø", "oe", -1)
		newStr = strings.Replace(newStr, "å", "aa", -1)
		res = append(res, newStr)
	}
	return res
}

func IntoTableQuoteType(headers []string) []TableQuoteType {
	types := TableQuoteTypesMap()
	res := []TableQuoteType{}
	for i := range headers {
		res = append(res, types[headers[i]])
	}
	return res
}

func TableQuoteTypesMap() map[string]TableQuoteType {
	return map[string]TableQuoteType{
		"id":                                     Quoted,   //VARCHAR(36) UNIQUE PRIMARY KEY NOT NULL",
		"status":                                 Unquoted, //bit NOT NULL",
		"oprettet":                               Quoted,   //timestamp NOT NULL",
		"aendret":                                Quoted,   //timestamp NOT NULL",
		"vejkode":                                Unquoted, //INT UNSIGNED NOT NULL ",
		"vejnavn":                                Quoted,   //text NOT NULL",
		"adresseringsvejnavn":                    Quoted,   //text NOT NULL",
		"husnr":                                  Quoted,   //text NOT NULL",
		"etage":                                  Quoted,   //text",
		"doer":                                   Quoted,   //text",
		"supplerendebynavn":                      Quoted,   //text",
		"postnr":                                 Unquoted, //INT UNSIGNED NOT NULL",
		"postnrnavn":                             Quoted,   //text NOT NULL",
		"stormodtagerpostnr":                     Unquoted, //INT UNSIGNED",
		"stormodtagerpostnrnavn":                 Quoted,   //text",
		"kommunekode":                            Unquoted, //INT UNSIGNED NOT NULL",
		"kommunenavn":                            Quoted,   //text NOT NULL",
		"ejerlavkode":                            Unquoted, //INT UNSIGNED NOT NULL",
		"ejerlavnavn":                            Quoted,   //text NOT NULL",
		"matrikelnr":                             Quoted,   //TEXT NOT NULL",
		"esrejendomsnr":                          Unquoted, //INT UNSIGNED NOT NULL",
		"etrs89koordinat_oest":                   Unquoted, //double NOT NULL",
		"etrs89koordinat_nord":                   Unquoted, //double NOT NULL",
		"wgs84koordinat_bredde":                  Unquoted, //double NOT NULL",
		"wgs84koordinat_laengde":                 Unquoted, //double NOT NULL",
		"noejagtighed":                           Quoted,   //VARCHAR(1) NOT NULL",
		"kilde":                                  Unquoted, //INT UNSIGNED NOT NULL",
		"tekniskstandard":                        Quoted,   //VARCHAR(2) NOT NULL",
		"tekstretning":                           Unquoted, //double NOT NULL",
		"ddkn_m100":                              Quoted,   //text NOT NULL",
		"ddkn_km1":                               Quoted,   //text NOT NULL",
		"ddkn_km10":                              Quoted,   //text NOT NULL",
		"adressepunktaendringsdato":              Quoted,   //timestamp NOT NULL",
		"adgangsadresseid":                       Quoted,   //VARCHAR(36) NOT NULL",
		"adgangsadresse_status":                  Unquoted, //bit NOT NULL",
		"adgangsadresse_oprettet":                Quoted,   //timestamp NOT NULL",
		"adgangsadresse_aendret":                 Quoted,   //timestamp NOT NULL",
		"regionskode":                            Unquoted, //INT UNSIGNED NOT NULL",
		"regionsnavn":                            Quoted,   //text NOT NULL",
		"jordstykke_ejerlavnavn":                 Quoted,   //text NOT NULL",
		"kvhx":                                   Quoted,   //text NOT NULL",
		"sognekode":                              Unquoted, //INT UNSIGNED NOT NULL",
		"sognenavn":                              Quoted,   //text NOT NULL",
		"politikredskode":                        Unquoted, //INT UNSIGNED NOT NULL",
		"politikredsnavn":                        Quoted,   //text NOT NULL",
		"retskredskode":                          Unquoted, //INT UNSIGNED NOT NULL",
		"retskredsnavn":                          Quoted,   //text NOT NULL",
		"opstillingskredskode":                   Unquoted, //INT UNSIGNED NOT NULL",
		"opstillingskredsnavn":                   Quoted,   //text NOT NULL",
		"zone":                                   Quoted,   //text NOT NULL",
		"jordstykke_ejerlavkode":                 Unquoted, //INT UNSIGNED NOT NULL",
		"jordstykke_matrikelnr":                  Quoted,   //text NOT NULL",
		"jordstykke_esrejendomsnr":               Unquoted, //INT UNSIGNED NOT NULL",
		"kvh":                                    Quoted,   //text NOT NULL",
		"hoejde":                                 Unquoted, //double NOT NULL",
		"adgangspunktid":                         Quoted,   //VARCHAR(36) NOT NULL",
		"vejpunkt_id":                            Quoted,   //VARCHAR(36) NOT NULL",
		"vejpunkt_kilde":                         Quoted,   //text NOT NULL",
		"vejpunkt_noejagtighed":                  Quoted,   //VARCHAR(1) NOT NULL",
		"vejpunkt_tekniskstandard":               Quoted,   //VARCHAR(2) NOT NULL",
		"vejpunkt_x":                             Unquoted, //double NOT NULL",
		"vejpunkt_y":                             Unquoted, //double NOT NULL",
		"afstemningsomraadenummer":               Unquoted, //INT UNSIGNED NOT NULL",
		"afstemningsomraadenavn":                 Quoted,   //text NOT NULL",
		"brofast":                                Unquoted, //bit NOT NULL",
		"supplerendebynavn_dagi_id":              Unquoted, //INT UNSIGNED",
		"navngivenvej_id":                        Quoted,   //VARCHAR(36) NOT NULL",
		"menighedsraadsafstemningsomraadenummer": Unquoted, //INT UNSIGNED NOT NULL",
		"menighedsraadsafstemningsomraadenavn":   Quoted,   //text NOT NULL",
		"vejpunkt_aendret":                       Quoted,   //timestamp NOT NULL",
		"ikrafttraedelse":                        Quoted,   //timestamp NOT NULL",
		"nedlagt":                                Unquoted, //bit",
		"adgangsadresse_ikrafttraedelse":         Quoted,   //timestamp NOT NULL",
		"adgangsadresse_nedlagt":                 Quoted,   //timestamp",
		"adgangsadresse_darstatus":               Unquoted, //INT UNSIGNED NOT NULL",
		"darstatus":                              Unquoted, //INT UNSIGNED NOT NULL",
		"storkredsnummer":                        Unquoted, //INT UNSIGNED NOT NULL",
		"storkredsnavn":                          Quoted,   //text NOT NULL",
		"valglandsdelsbogstav":                   Quoted,   //VARCHAR(1) NOT NULL",
		"valglandsdelsnavn":                      Quoted,   //text NOT NULL",
		"landsdelsnuts3":                         Quoted,   //text NOT NULL",
		"landsdelsnavn":                          Quoted,   //text NOT NULL",
		"betegnelse":                             Quoted,   //text NOT NULL",
	}
}

func TableSQLTypes() map[string]string {
	return map[string]string{
		"id":                                     "VARCHAR(36) UNIQUE PRIMARY KEY NOT NULL",
		"status":                                 "TINYINT UNSIGNED NOT NULL",
		"oprettet":                               "timestamp NOT NULL",
		"aendret":                                "timestamp NOT NULL",
		"vejkode":                                "SMALLINT UNSIGNED NOT NULL ",
		"vejnavn":                                "text NOT NULL",
		"adresseringsvejnavn":                    "text NOT NULL",
		"husnr":                                  "text NOT NULL",
		"etage":                                  "text",
		"doer":                                   "text",
		"supplerendebynavn":                      "text",
		"postnr":                                 "SMALLINT UNSIGNED NOT NULL",
		"postnrnavn":                             "text NOT NULL",
		"stormodtagerpostnr":                     "TINYINT UNSIGNED",
		"stormodtagerpostnrnavn":                 "text",
		"kommunekode":                            "SMALLINT UNSIGNED NOT NULL",
		"kommunenavn":                            "text NOT NULL",
		"ejerlavkode":                            "INT UNSIGNED",
		"ejerlavnavn":                            "text",
		"matrikelnr":                             "TEXT",
		"esrejendomsnr":                          "INT UNSIGNED",
		"etrs89koordinat_oest":                   "double NOT NULL",
		"etrs89koordinat_nord":                   "double NOT NULL",
		"wgs84koordinat_bredde":                  "double NOT NULL",
		"wgs84koordinat_laengde":                 "double NOT NULL",
		"noejagtighed":                           "VARCHAR(1) NOT NULL",
		"kilde":                                  "TINYINT UNSIGNED",
		"tekniskstandard":                        "VARCHAR(2) NOT NULL",
		"tekstretning":                           "double NOT NULL",
		"ddkn_m100":                              "text NOT NULL",
		"ddkn_km1":                               "text NOT NULL",
		"ddkn_km10":                              "text NOT NULL",
		"adressepunktaendringsdato":              "timestamp NOT NULL",
		"adgangsadresseid":                       "VARCHAR(36) NOT NULL",
		"adgangsadresse_status":                  "TINYINT UNSIGNED NOT NULL",
		"adgangsadresse_oprettet":                "timestamp NOT NULL",
		"adgangsadresse_aendret":                 "timestamp NOT NULL",
		"regionskode":                            "SMALLINT UNSIGNED NOT NULL",
		"regionsnavn":                            "text NOT NULL",
		"jordstykke_ejerlavnavn":                 "text",
		"kvhx":                                   "text NOT NULL",
		"sognekode":                              "SMALLINT UNSIGNED NOT NULL",
		"sognenavn":                              "text NOT NULL",
		"politikredskode":                        "SMALLINT UNSIGNED NOT NULL",
		"politikredsnavn":                        "text NOT NULL",
		"retskredskode":                          "SMALLINT UNSIGNED NOT NULL",
		"retskredsnavn":                          "text NOT NULL",
		"opstillingskredskode":                   "TINYINT UNSIGNED NOT NULL",
		"opstillingskredsnavn":                   "text NOT NULL",
		"zone":                                   "text NOT NULL",
		"jordstykke_ejerlavkode":                 "INT UNSIGNED",
		"jordstykke_matrikelnr":                  "text",
		"jordstykke_esrejendomsnr":               "INT UNSIGNED",
		"kvh":                                    "text NOT NULL",
		"hoejde":                                 "double",
		"adgangspunktid":                         "VARCHAR(36) NOT NULL",
		"vejpunkt_id":                            "VARCHAR(36) NOT NULL",
		"vejpunkt_kilde":                         "text NOT NULL",
		"vejpunkt_noejagtighed":                  "VARCHAR(1) NOT NULL",
		"vejpunkt_tekniskstandard":               "VARCHAR(2) NOT NULL",
		"vejpunkt_x":                             "double NOT NULL",
		"vejpunkt_y":                             "double NOT NULL",
		"afstemningsomraadenummer":               "TINYINT UNSIGNED NOT NULL",
		"afstemningsomraadenavn":                 "text NOT NULL",
		"brofast":                                "bit NOT NULL",
		"supplerendebynavn_dagi_id":              "INT UNSIGNED",
		"navngivenvej_id":                        "VARCHAR(36) NOT NULL",
		"menighedsraadsafstemningsomraadenummer": "TINYINT UNSIGNED NOT NULL",
		"menighedsraadsafstemningsomraadenavn":   "text NOT NULL",
		"vejpunkt_aendret":                       "timestamp NOT NULL",
		"ikrafttraedelse":                        "timestamp",
		"nedlagt":                                "bit",
		"adgangsadresse_ikrafttraedelse":         "timestamp",
		"adgangsadresse_nedlagt":                 "timestamp",
		"adgangsadresse_darstatus":               "TINYINT UNSIGNED NOT NULL",
		"darstatus":                              "TINYINT UNSIGNED NOT NULL",
		"storkredsnummer":                        "TINYINT UNSIGNED NOT NULL",
		"storkredsnavn":                          "text NOT NULL",
		"valglandsdelsbogstav":                   "VARCHAR(1) NOT NULL",
		"valglandsdelsnavn":                      "text NOT NULL",
		"landsdelsnuts3":                         "text NOT NULL",
		"landsdelsnavn":                          "text NOT NULL",
		"betegnelse":                             "text NOT NULL",
	}
}