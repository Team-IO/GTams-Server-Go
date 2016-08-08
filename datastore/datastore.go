package datastore

import(
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"./../core"
	"./entities"
	"github.com/satori/go.uuid"
)

type Datastore struct {

}

func InitDatastore() {
	core.Logger.Info("Initializing Datastore...")


}

func GetInstallation(id uuid.UUID) {

}

func NewInstallation(installation entities.Installation) entities.Installation {
	return entities.Installation{
		Id: uuid.NewV4(),
	}
}

func NewTerminal(owner uuid.UUID) entities.Terminal {
	return entities.Terminal {

	}
}