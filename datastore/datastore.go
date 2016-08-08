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

var insertInstallation *sql.Stmt

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

func prepareStatements() {
	var err error
	insertInstallation, err = db.Prepare("INSERT INTO installation (uuid, version, mcversion, branding, language) VALUES (?, ?, ?, ?, ?);")

	if err != nil {
		panic(err.Error())
	}
}

func CloseDatastore() {
	if insertInstallation {
		insertInstallation.Close()
	}

	db.Close()
}

func GetInstallation(id uuid.UUID) {
}

func saveInstallation(installation entities.Installation) {
	result, err := insertInstallation.Exec(installation.Id, installation.Version, installation.McVersion, installation.Branding, installation.Language)

}

func NewInstallation(installation entities.Installation) entities.Installation {
	return entities.Installation{
		Id: uuid.NewV4(),
	}
}

func NewTerminal(owner uuid.UUID) entities.Terminal {
	return entities.Terminal{

	}
}