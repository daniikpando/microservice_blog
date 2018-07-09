package connections

import (
	"github.com/jinzhu/gorm"
	"sync"
	"fmt"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/daniel/basic_microservice/article/configuration"
	"github.com/labstack/gommon/log"
)

var (
	connect sync.Once
	db      *gorm.DB
	err error = nil
)

// GetConnection se encarga de abrir la conexión a la base de datos POSTGRES manejando el patrón singleton y generando un pool de conexiones
func GetConnection() (*gorm.DB, error) {
	connect.Do(func() {
		var config configuration.Configuration

		config, err = configuration.GetConfiguration()

		if err != nil {
			return
		}

		stringConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			config.Server,
			config.User,
			config.Database,
			config.Password,
		)

		db, err = gorm.Open(config.Engine, stringConnection)
	})

	if err == nil{
		err = db.DB().Ping()
	}else{
		log.Error(err)
	}

	return db, err
}
