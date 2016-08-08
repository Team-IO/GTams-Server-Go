package datastore

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"./../core"
	"./entities"
	"github.com/satori/go.uuid"
)

var db *sql.DB
var installations map[uuid.UUID]entities.Installation

var statements [20]*sql.Stmt

var insertInstallation *sql.Stmt
var updateInstallation *sql.Stmt
var selectInstallation *sql.Stmt
var selectInstallationUUID *sql.Stmt

func InitDatastore() {
	core.Logger.Info("Initializing Datastore...")
	var err error
	db, err = sql.Open("mysql", "gtams:gtams@/GTams")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	prepareStatements()
}

func CloseDatastore() {
	for _,statement := range statements {
		err := statement.Close()
		if err != nil {
			core.Logger.Error("Error closing prepared statement: %v", err.Error())
		}
	}

	err := db.Close()
	if err != nil {
		core.Logger.Error("Error closing database: %v", err.Error())
	}
}

func statement(query string) *sql.Stmt {
	var statement, err = db.Prepare(query)

	if err != nil {
		panic(err.Error())
	}
	statements = append(statements, statement)
	return statement
}

func prepareStatements() {
	insertInstallation = statement("INSERT INTO installation (uuid, version, mcversion, branding, language) VALUES (?, ?, ?, ?, ?);")
	updateInstallation = statement("UPDATE installation SET version=?, mcversion=?, branding=?, language=? WHERE uuid=?;")
	selectInstallation = statement("SELECT last_seen, version, mcversion, branding, language FROM installation WHERE uuid=?")
	selectInstallationUUID = statement("SELECT uuid FROM installation WHERE uuid=?")
}

func GetInstallation(id uuid.UUID, allFields bool) entities.Installation {
	installation := installations[id]
	if installation == nil {
		var stmt *sql.Stmt
		if allFields {
			stmt = selectInstallation
		} else {
			stmt = selectInstallationUUID
		}
		rows, err := stmt.Query(id.String())
		if err != nil {
			core.Logger.Error("Error reading installation from DB: %v", err.Error())
			panic(err.Error())
		}
		defer rows.Close()
		if rows.Next() {
			installation = entities.Installation{
				Id: id,
			}
			if allFields {
				err = rows.Scan(&installation.LastSeen, &installation.Version, &installation.McVersion, &installation.Branding, & installation.Language)
			}
			// Yes, this returns an installation entity with only an id -> not nil as in "installation exists"
			return installation
		}
	}
	return nil
}

func insertInstallation(installation entities.Installation) {
	_, err := insertInstallation.Exec(installation.Id, installation.Version, installation.McVersion, installation.Branding, installation.Language)

	if err != nil {
		core.Logger.Error("Error inserting installation in DB: %v", err.Error())
		panic(err.Error())
	}
}

func updateInstallation(installation entities.Installation) {
	_, err := updateInstallation.Exec(installation.Version, installation.McVersion, installation.Branding, installation.Language, installation.Id)

	if err != nil {
		core.Logger.Error("Error updating installation in DB: %v", err.Error())
		panic(err.Error())
	}
}

// If installation information is given, it will be checked if that info is in the DB & will update the timestamp.
// If no information is given or the info is not yet in the db a new id will be generated & stored in DB.
func PingInstallation(installation entities.Installation) uuid.UUID {
	if installation == nil {
		// New installation with blank values & new ID
		installation = entities.Installation {
			Id: uuid.NewV4(),
		}
		insertInstallation(installation)
	} else {
		inDatabase := GetInstallation(installation.Id, false)
		if inDatabase == nil {
			installation.Id = uuid.NewV4()
			insertInstallation(installation)
		} else {
			// Always save, if only to update the timestamp
			updateInstallation(installation)
		}
	}
	return installation.Id
}

func NewTerminal(owner uuid.UUID) entities.Terminal {
	return entities.Terminal{

	}
}