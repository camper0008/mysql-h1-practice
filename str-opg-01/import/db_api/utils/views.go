package utils

import (
	"csv-to-mysql/table/defs"
	"csv-to-mysql/utils"
	"database/sql"
	"fmt"
)

func Postcard(db *sql.DB) {
	_, err := db.Exec("DROP VIEW IF EXISTS postkort")
	utils.LogError(err)

	selectQuery := (`
		SELECT 
			adresse.id, adresse.vej_navn, 
			adresse.hus_nr, adresse.etage, 
			adresse.doer, adresse.post_nr, 
			post_nr_map.navn AS post_nr_navn, 
			region_nr_map.navn AS region
		FROM (adresse)
		JOIN (post_nr_map) ON (adresse.post_nr = post_nr_map.nr)
		JOIN (region_nr_map) ON (adresse.region_nr = region_nr_map.nr)
	`)
	createQuery := fmt.Sprintf("CREATE VIEW postkort AS %s", selectQuery)
	_, err = db.Exec(createQuery)
	utils.LogError(err)
}

func Address(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS adresse")
	utils.LogError(err)

	types := defs.TableSQLTypes()
	createQuery := fmt.Sprintf(`
		CREATE TABLE adresse (
			id %s,
			vej_navn %s,
			hus_nr %s,
			etage %s,
			doer %s,
			post_nr %s,
			region_nr %s,
			FOREIGN KEY (post_nr) REFERENCES post_nr_map(nr),
			FOREIGN KEY (region_nr) REFERENCES region_nr_map(nr)
		)`,
		types["id"],
		types["vejnavn"],
		types["husnr"],
		types["etage"],
		types["doer"],
		types["postnr"],
		types["regionskode"])

	_, err = db.Exec(createQuery)
	utils.LogError(err)

	insertQuery := (`
		INSERT INTO adresse (
			SELECT id, vejnavn AS vej_navn, husnr AS hus_nr, etage, doer, postnr AS post_nr, regionskode AS region_nr FROM raw
		)`)
	_, err = db.Exec(insertQuery)
	utils.LogError(err)
}

func PostCodeMap(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS adresse, post_nr_map")
	utils.LogError(err)

	types := defs.TableSQLTypes()
	createQuery := fmt.Sprintf(`CREATE TABLE post_nr_map (nr %s, navn %s)`,
		types["postnr"]+" UNIQUE PRIMARY KEY",
		types["postnrnavn"])

	_, err = db.Exec(createQuery)
	utils.LogError(err)

	valuesQuery := `SELECT DISTINCT postnr AS nr, postnrnavn AS navn FROM raw`
	insertQuery := fmt.Sprintf("INSERT INTO post_nr_map %s", valuesQuery)
	_, err = db.Exec(insertQuery)
	utils.LogError(err)
}

func RegionCodeMap(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS adresse, region_nr_map")
	utils.LogError(err)

	types := defs.TableSQLTypes()
	createQuery := fmt.Sprintf(`CREATE TABLE region_nr_map (nr %s, navn %s)`,
		types["regionskode"]+" UNIQUE PRIMARY KEY",
		types["regionsnavn"])

	_, err = db.Exec(createQuery)
	utils.LogError(err)

	valuesQuery := `SELECT DISTINCT regionskode AS nr, regionsnavn AS navn FROM raw`
	insertQuery := fmt.Sprintf("INSERT INTO region_nr_map %s", valuesQuery)
	_, err = db.Exec(insertQuery)
	utils.LogError(err)
}
