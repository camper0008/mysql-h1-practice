package utils

import (
	"csv-to-mysql/table/defs"
	"csv-to-mysql/utils"
	"database/sql"
	"fmt"
)

func SuppliedCityName(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS supplerende_by")
	utils.LogError(err)

	types := defs.TableSQLTypes()
	createQuery := fmt.Sprintf(`
		CREATE TABLE supplerende_by (
			id %s,
			navn %s
		)`,
		types["id"],
		types["supplerendebynavn"],
	)

	_, err = db.Exec(createQuery)
	utils.LogError(err)

	insertQuery := (`
		INSERT INTO supplerende_by (
			SELECT UUID() as id, unik_suplerende_by.navn FROM 
			(SELECT DISTINCT raw.supplerendebynavn AS navn FROM raw) AS unik_suplerende_by
		)`)
	_, err = db.Exec(insertQuery)
	utils.LogError(err)
}

func Postcard(db *sql.DB) {
	_, err := db.Exec("DROP VIEW IF EXISTS postkort")
	utils.LogError(err)

	selectQuery := (`
		SELECT 
			vej.navn as vej_navn, 
			hus.nr, 
			hus.etage, 
			hus.doer,
			supplerende_by.navn as supplerende_by_navn, 
			vej.post_nr, 
			post_nr_map.navn AS post_nr_navn, 
			region_nr_map.navn AS region
		FROM (hus)
		JOIN (supplerende_by) ON (hus.supplerende_by_id = supplerende_by.id)
		JOIN (vej) ON (vej.id = hus.vej_id)
		JOIN (post_nr_map) ON (vej.post_nr = post_nr_map.nr)
		JOIN (region_nr_map) ON (vej.region_nr = region_nr_map.nr)
	`)
	createQuery := fmt.Sprintf("CREATE VIEW postkort AS %s", selectQuery)
	_, err = db.Exec(createQuery)
	utils.LogError(err)
}

func House(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS hus")
	utils.LogError(err)

	types := defs.TableSQLTypes()
	createQuery := fmt.Sprintf(`
		CREATE TABLE hus (
			id %s,
			nr %s,
			etage %s,
			doer %s,
			supplerende_by_id %s,
			vej_id %s,
			FOREIGN KEY (supplerende_by_id) REFERENCES supplerende_by(id),
			FOREIGN KEY (vej_id) REFERENCES vej(id)
		)`,
		types["id"],
		types["husnr"],
		types["etage"],
		types["doer"],
		"VARCHAR(36) NOT NULL",
		"VARCHAR(36) NOT NULL")

	_, err = db.Exec(createQuery)
	utils.LogError(err)

	insertQuery := (`
		INSERT INTO hus (
			SELECT 
				UUID() AS id, 
				raw.husnr AS nr,
				raw.etage,
				raw.doer,
				supplerende_by.id as supplerende_by_id,
				vej.id as vej_id
			FROM raw
			JOIN vej
			ON (raw.vejnavn = vej.navn) AND (raw.postnr = vej.post_nr)
			JOIN supplerende_by
			ON (raw.supplerendebynavn = supplerende_by.navn)
		)`)

	_, err = db.Exec(insertQuery)
	utils.LogError(err)
}

func Road(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS hus, vej")
	utils.LogError(err)

	types := defs.TableSQLTypes()
	createQuery := fmt.Sprintf(`
		CREATE TABLE vej (
			id %s,
			navn %s,
			post_nr %s,
			region_nr %s,
			FOREIGN KEY (post_nr) REFERENCES post_nr_map(nr),
			FOREIGN KEY (region_nr) REFERENCES region_nr_map(nr)
		)`,
		types["id"],
		types["vejnavn"],
		types["postnr"],
		types["regionskode"],
	)

	_, err = db.Exec(createQuery)
	utils.LogError(err)

	insertQuery := (`
		INSERT INTO vej (
			SELECT UUID() as id, vej.navn, vej.post_nr, vej.region_nr FROM 
			(SELECT DISTINCT vejnavn AS navn, postnr AS post_nr, regionskode AS region_nr FROM raw) AS vej
		)`)
	_, err = db.Exec(insertQuery)
	utils.LogError(err)
}

func PostCodeMap(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS hus, vej, post_nr_map")
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
	_, err := db.Exec("DROP TABLE IF EXISTS hus, vej, region_nr_map")
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
