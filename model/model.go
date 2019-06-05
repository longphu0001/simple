package model

import (
	"go-learn/projects/simple/config"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
)

// Service is the interface of all model service.
type Service interface {
}

type service struct {
}

// New returns a Service instance for operating all model service.
func New(dbCfg *config.Database) (Service, error) {
	_, err := newDB(dbCfg) //needs to pass db as
	if err != nil {
		return nil, errors.Wrap(err, "Failed to initialize db of grom")
	}

	serv := &service{}
	return serv, nil
}

func newDB(dbCfg *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbCfg.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db URL")
	}
	db.DB().SetMaxOpenConns(dbCfg.MaxActive)
	db.DB().SetMaxIdleConns(dbCfg.MaxIdle)

	db.LogMode(dbCfg.LogMode)
	return db, nil
}
